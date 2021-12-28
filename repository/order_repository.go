package repository

import (
	"github.com/edwardsuwirya/wmbPos/entity"
	"gorm.io/gorm"
	"log"
)

type IOrderRepository interface {
	CreateOne(order entity.CustomerOrder) (*entity.CustomerOrder, error)
	UpdatePaymentMethod(orderPayment entity.OrderPayment) (string, error)
	GetSummaryPrice(billNo string) (int, error)
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

func (o *OrderRepository) UpdatePaymentMethod(orderPayment entity.OrderPayment) (string, error) {
	result := o.db.Create(orderPayment)
	if result.Error != nil {
		log.Println(result.Error)
		return "", result.Error
	}
	return orderPayment.CustomerOrderID, nil
}

func (o *OrderRepository) GetSummaryPrice(billNo string) (int, error) {
	var total int
	result := o.db.Model(&entity.CustomerOrderDetail{}).
		Select("sum(qty * price)").
		Where("customer_order_id = ?", billNo).
		Scan(&total)
	if result.Error != nil {
		log.Println(result.Error)
		return -1, result.Error
	}
	return total, nil
}

func NewOrderRepository(resource *gorm.DB) IOrderRepository {
	menuRepo := &OrderRepository{db: resource}
	return menuRepo
}
