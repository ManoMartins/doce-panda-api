package dtos

import "doce-panda/domain/coupon/entity"

type InputRequestExchangeOrderDto struct {
	ID         string `json:"id"`
	UserID     string `json:"userId"`
	OrderItems []struct {
		ProductID string `json:"productId"`
		Quantity  int    `json:"quantity"`
	} `json:"orderItems"`
}

type OutputRequestExchangeOrderDto struct {
	CouponID    string            `json:"couponId"`
	VoucherCode string            `json:"voucherCode"`
	Amount      int               `json:"amount"`
	Status      entity.StatusEnum `json:"status"`
}
