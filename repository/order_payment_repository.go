package repository

import (
	"github.com/edwardsuwirya/wmbPos/entity"
	"gorm.io/gorm"
	"log"
)

type IOrderPaymentRepository interface {
	CreateOne(orderPayment entity.OrderPayment) (*entity.OrderPayment, error)
	GetPaymentMethodById(id string) (*entity.Payment, error)
}

type OrderPaymentRepository struct {
	db *gorm.DB
}

func (o *OrderPaymentRepository) CreateOne(orderPayment entity.OrderPayment) (*entity.OrderPayment, error) {
	err := o.db.Create(&orderPayment).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &orderPayment, nil
}

func (o *OrderPaymentRepository) GetPaymentMethodById(id string) (*entity.Payment, error) {
	var payment entity.Payment
	err := o.db.Where("id = ?", id).First(&payment).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &payment, nil
}
func NewOrderPaymentRepository(resource *gorm.DB) IOrderPaymentRepository {
	paymentRepo := &OrderPaymentRepository{db: resource}
	return paymentRepo
}
