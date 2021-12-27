package usecase

import (
	"github.com/edwardsuwirya/wmbPos/apperror"
	"github.com/edwardsuwirya/wmbPos/dto"
	"github.com/edwardsuwirya/wmbPos/entity"
	"github.com/edwardsuwirya/wmbPos/repository"
	"gorm.io/gorm"
)

type IOrderUseCase interface {
	OpenOrder(order dto.CustomerOrderRequest) (*entity.CustomerOrder, error)
	CloseOrder(billNo string, paymentMethod string) (string, error)
}

type OrderUseCase struct {
	orderRepo   repository.IOrderRepository
	resvUseCase IOrderTableReservationUseCase
}

func NewOrderUseCase(orderRepo repository.IOrderRepository, resvUseCase IOrderTableReservationUseCase) IOrderUseCase {
	return &OrderUseCase{
		orderRepo:   orderRepo,
		resvUseCase: resvUseCase,
	}
}

func (o *OrderUseCase) OpenOrder(order dto.CustomerOrderRequest) (*entity.CustomerOrder, error) {
	var orderDetails []entity.CustomerOrderDetail
	for _, od := range order.Orders {
		orderDetail := entity.CustomerOrderDetail{
			MenuID: od.MenuId,
			Qty:    od.Qty,
		}
		orderDetails = append(orderDetails, orderDetail)
	}
	ord, err := o.orderRepo.CreateOne(entity.CustomerOrder{
		CustomerName:  order.CustomerName,
		PaymentMethod: "",
		OrderDetails:  orderDetails,
		Model:         gorm.Model{},
	})
	if err != nil {
		return nil, err
	}
	err = o.resvUseCase.ReserveTable(dto.TableRequest{
		BillNo:  ord.ID,
		TableId: order.TableId,
	})
	if err != nil {
		return nil, apperror.TableOccupiedError
	}
	return ord, nil
}

func (o *OrderUseCase) CloseOrder(billNo string, paymentMethod string) (string, error) {
	billId, err := o.orderRepo.UpdatePaymentMethod(billNo, paymentMethod)
	if err != nil {
		return "", err
	}
	err = o.resvUseCase.CloseTable(billId)
	if err != nil {
		return billId, err
	}
	return billId, nil
}
