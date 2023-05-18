package entity

import (
	couponEntity "doce-panda/domain/coupon/entity"
	"doce-panda/domain/payment/entity"
	userEntity "doce-panda/domain/user/entity"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"math"
	"time"
)

type StatusEnum string

const (
	WAITING_PAYMENT         StatusEnum = "WAITING_PAYMENT"
	PREPARING               StatusEnum = "PREPARING"
	IN_TRANSIT              StatusEnum = "IN_TRANSIT"
	DELIVERED               StatusEnum = "DELIVERED"
	EXCHANGE_REQUEST        StatusEnum = "EXCHANGE_REQUEST"
	ACCEPT_EXCHANGE_REQUEST StatusEnum = "ACCEPT_EXCHANGE_REQUEST"
	EXCHANGE_RECEIVED       StatusEnum = "EXCHANGE_RECEIVED"
	DENY_EXCHANGE_REQUEST   StatusEnum = "DENY_EXCHANGE_REQUEST"
)

type Order struct {
	ID              string                `json:"id" validate:"required"`
	OrderItems      []OrderItem           `json:"orderItems"`
	SubTotalInCents int                   `json:"subTotalInCents"`
	TotalInCents    int                   `json:"totalInCents"`
	Status          StatusEnum            `json:"status" validate:"required,oneof='WAITING_PAYMENT' 'PREPARING' 'IN_TRANSIT' 'DELIVERED' 'EXCHANGE_REQUEST' 'EXCHANGE_RECEIVED' 'ACCEPT_EXCHANGE_REQUEST' 'DENY_EXCHANGE_REQUEST'"`
	Payments        []entity.CreditCard   `json:"payments"`
	DeliveredFee    int                   `json:"deliveredFee"`
	CouponID        *string               `json:"couponId"`
	Coupons         []couponEntity.Coupon `json:"coupons"`
	AddressID       string                `json:"addressId"`
	Address         *userEntity.Address   `json:"address"`
	UserID          string                `json:"userId"`
	User            *userEntity.User      `json:"user"`
	CreatedAt       time.Time             `json:"createdAt"`
	UpdatedAt       time.Time             `json:"updatedAt"`
}

type OrderInterface interface {
	Validate() error
	AddOrderItems(orderItems []OrderItem)
	UpdateStatus(status StatusEnum) error
	RequestExchange() error
	ExchangeReceived() error
	ApplyCoupon(coupon couponEntity.Coupon) (int, error)
	ApplyCoupons(coupons []couponEntity.Coupon) (int, error)
	AcceptExchangeRequest() error
	DenyExchangeRequest() error
}

func NewOrder(order Order) (*Order, error) {
	o := Order{
		ID:              order.ID,
		OrderItems:      order.OrderItems,
		SubTotalInCents: order.SubTotalInCents,
		TotalInCents:    order.TotalInCents,
		Status:          order.Status,
		Payments:        order.Payments,
		DeliveredFee:    order.DeliveredFee,
		CouponID:        order.CouponID,
		Coupons:         order.Coupons,
		AddressID:       order.AddressID,
		UserID:          order.UserID,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       order.UpdatedAt,
	}

	if order.ID == "" {
		o.ID = uuid.NewString()
	}

	if order.Status == "" {
		o.Status = WAITING_PAYMENT
	}

	if order.CouponID == nil && order.TotalInCents == 0 {
		o.TotalInCents = o.SubTotalInCents + o.DeliveredFee
	}

	if order.Address != nil {
		o.Address = order.Address
	}

	if order.User != nil {
		o.User = &userEntity.User{
			ID:             order.User.ID,
			Name:           order.User.Name,
			Gender:         order.User.Gender,
			PhoneNumber:    order.User.PhoneNumber,
			DocumentNumber: order.User.DocumentNumber,
			RewardPoints:   order.User.RewardPoints,
			Email:          order.User.Email,
			CreatedAt:      order.User.CreatedAt,
			UpdatedAt:      order.User.UpdatedAt,
		}
	}

	err := o.Validate()

	if err != nil {
		return nil, err
	}

	return &o, nil
}

func (o *Order) Validate() error {
	var Validator = validator.New()

	return Validator.Struct(*o)
}

func (o *Order) AddOrderItems(orderItems []OrderItem) {
	for i := range orderItems {
		orderItems[i].OrderID = o.ID
	}

	o.OrderItems = orderItems
}

func (o *Order) UpdateStatus(status StatusEnum) error {
	o.Status = status

	return o.Validate()
}

func (o *Order) RequestExchange() error {
	if o.Status != DELIVERED {
		return fmt.Errorf("O pedido precisa está com status de entregue")
	}

	o.Status = EXCHANGE_REQUEST

	return o.Validate()
}

func (o *Order) ApplyCoupon(coupon couponEntity.Coupon) (int, error) {
	if coupon.Status == couponEntity.USED {
		return 0, fmt.Errorf("O cupom já foi utilizado")
	}

	if coupon.UserID != o.UserID {
		return 0, fmt.Errorf("O cupom não pertence ao usuário")
	}

	o.TotalInCents = o.SubTotalInCents + o.DeliveredFee - coupon.Amount

	if o.TotalInCents < 0 {
		moneyExchange := o.TotalInCents
		o.TotalInCents = 0
		return int(math.Abs(float64(moneyExchange))), nil
	}

	return 0, nil
}

func (o *Order) ApplyCoupons(coupons []couponEntity.Coupon) (int, error) {
	couponAmount := 0
	for _, coupon := range coupons {
		if coupon.Status == couponEntity.USED {
			return 0, fmt.Errorf("O cupom %s  já foi utilizado", coupon.VoucherCode)
		}

		if coupon.UserID != o.UserID {
			return 0, fmt.Errorf("O cupom %s não pertence ao usuário", coupon.VoucherCode)
		}

		couponAmount += coupon.Amount
	}

	o.TotalInCents = o.SubTotalInCents + o.DeliveredFee - couponAmount

	if o.TotalInCents < 0 {
		moneyExchange := o.TotalInCents
		o.TotalInCents = 0
		return int(math.Abs(float64(moneyExchange))), nil
	}

	return 0, nil
}

func (o *Order) AcceptExchangeRequest() error {
	if o.Status != EXCHANGE_REQUEST {
		return fmt.Errorf("Deve ter sido solicitado a troca")
	}

	if o.Status == DENY_EXCHANGE_REQUEST {
		return fmt.Errorf("O pedido já teve a troca negada")
	}

	o.Status = ACCEPT_EXCHANGE_REQUEST

	return nil
}

func (o *Order) DenyExchangeRequest() error {
	if o.Status != EXCHANGE_REQUEST {
		return fmt.Errorf("Deve ter sido solicitado a troca")
	}

	if o.Status == ACCEPT_EXCHANGE_REQUEST {
		return fmt.Errorf("O pedido já teve a troca aceita")
	}

	o.Status = DENY_EXCHANGE_REQUEST

	return nil
}

func (o *Order) ExchangeReceived() error {
	if o.Status != ACCEPT_EXCHANGE_REQUEST {
		return fmt.Errorf("A troca do pedido precisa está aceito")
	}

	o.Status = EXCHANGE_RECEIVED

	return nil
}
