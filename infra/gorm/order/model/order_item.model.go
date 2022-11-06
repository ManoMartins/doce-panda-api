package model

import "doce-panda/infra/gorm/product/model"

type OrderItem struct {
	ID           string `json:"id" gorm:"type:uuid;primary_key"`
	ProductID    string `json:"productId" gorm:"column:product_id;type:uuid;notnull"`
	OrderID      string `json:"orderId" gorm:"column:order_id;type:uuid;notnull"`
	Quantity     int    `json:"quantity" gorm:"type:integer"`
	TotalInCents int    `json:"totalInCents" gorm:"type:integer"`
	Product      model.Product
}
