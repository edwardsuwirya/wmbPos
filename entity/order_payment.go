package entity

import "gorm.io/gorm"

type OrderPayment struct {
	CustomerOrderID string
	BillerReceipt   string `gorm:"size:50;"`
	PaymentID       string
	Payment         Payment
	gorm.Model
}

func (c *OrderPayment) TableName() string {
	return "order_payment"
}
