package entity

type Payment struct {
	ID                string `gorm:"column:id;size:3;primaryKey"`
	PaymentMethodName string `gorm:"size:20;"`
}

func (c *Payment) TableName() string {
	return "payment"
}
