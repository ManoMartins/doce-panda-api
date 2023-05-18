package dtos

import (
	couponEntity "doce-panda/domain/coupon/entity"
	"doce-panda/domain/order/entity"
)

type InputCreateOrderDto struct {
	AddressID  string `json:"addressId"`
	UserID     string `json:"userId"`
	OrderItems []struct {
		ProductID string `json:"productId"`
		Quantity  int    `json:"quantity"`
	} `json:"orderItems"`
	Payments []struct {
		PaymentID    string `json:"paymentId"`
		TotalInCents int    `json:"totalInCents"`
	} `json:"payments"`
	VoucherCode  string   `json:"voucherCode"`
	VoucherCodes []string `json:"voucherCodes"`
}

type OutputCreateOrderDto struct {
	ID                  string               `json:"id"`
	OrderItems          []entity.OrderItem   `json:"orderItems"`
	TotalInCents        int                  `json:"totalInCents"`
	Status              entity.StatusEnum    `json:"status"`
	CouponMoneyExchange *couponEntity.Coupon `json:"couponMoneyExchange"`
}
