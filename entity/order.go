package entity

import (
	guuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerOrder struct {
	ID           string `gorm:"column:id;size:36;primaryKey"`
	CustomerName string `gorm:"size:36;"`
	OrderPayment OrderPayment
	OrderDetails []CustomerOrderDetail
	gorm.Model
}

func (c *CustomerOrder) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = guuid.New().String()
	return nil
}
func (c *CustomerOrder) TableName() string {
	return "customer_order"
}
