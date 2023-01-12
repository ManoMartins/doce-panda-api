package model

import (
	"doce-panda/domain/coupon/entity"
	"doce-panda/infra/gorm/user/model"
	"time"
)

type Coupon struct {
	ID          string            `json:"id" gorm:"type:uuid;primary_key"`
	VoucherCode string            `json:"voucherCode" gorm:"type:varchar(255)"`
	Status      entity.StatusEnum `json:"status" gorm:"type:varchar(255)"`
	User        model.User        `json:"user" gorm:"ForeignKey:UserID"`
	UserID      string            `json:"userId" gorm:"column:user_id;type:uuid;notnull"`
	Amount      int               `json:"amount"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
}
