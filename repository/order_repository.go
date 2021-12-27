package repository

import (
	"errors"
	"github.com/edwardsuwirya/wmbPos/entity"
	"gorm.io/gorm"
	"log"
)

type IOrderRepository interface {
	CreateOne(order entity.CustomerOrder) (*entity.CustomerOrder, error)
	UpdatePaymentMethod(billNo string, payment string) (string, error)
}

type OrderRepository struct {
	db *gorm.DB
}

func (o *OrderRepository) CreateOne(order entity.CustomerOrder) (*entity.CustomerOrder, error) {
	err := o.db.Create(&order).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &order, nil
}

func (o *OrderRepository) UpdatePaymentMethod(billNo string, payment string) (string, error) {
	result := o.db.Model(&entity.CustomerOrder{}).Where("id = ?", billNo).Update("payment_method", payment)
	if result.Error != nil {
		log.Println(result.Error)
		return "", result.Error
	}
	if result.RowsAffected == 0 {
		return "", errors.New("Bill No Unrecognized")
	}
	return billNo, nil
}

func NewOrderRepository(resource *gorm.DB) IOrderRepository {
	menuRepo := &OrderRepository{db: resource}
	return menuRepo
}
