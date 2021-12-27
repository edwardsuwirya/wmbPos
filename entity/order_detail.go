package entity

import "gorm.io/gorm"

type CustomerOrderDetail struct {
	CustomerOrderID string
	MenuID  string
	Qty     int
	gorm.Model
}

func (c *CustomerOrderDetail) TableName() string {
	return "customer_order_detail"
}
