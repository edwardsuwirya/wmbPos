package usecase

import (
	"github.com/edwardsuwirya/wmbPos/apperror"
	"github.com/edwardsuwirya/wmbPos/dto"
	"github.com/edwardsuwirya/wmbPos/entity"
	"github.com/edwardsuwirya/wmbPos/repository"
	"gorm.io/gorm"
	"log"
)

type IOrderUseCase interface {
	OpenOrder(order dto.CustomerOrderRequest) (*entity.CustomerOrder, error)
	CloseOrder(closeOrderInfo dto.CloseOrderRequest) (string, error)
}

type OrderUseCase struct {
	orderRepo           repository.IOrderRepository
	resvUseCase         IOrderTableReservationUseCase
	orderPaymentUseCase IOrderPaymentUseCase
}

func NewOrderUseCase(orderRepo repository.IOrderRepository, resvUseCase IOrderTableReservationUseCase, paymentUseCase IOrderPaymentUseCase) IOrderUseCase {
	return &OrderUseCase{
		orderRepo:           orderRepo,
		resvUseCase:         resvUseCase,
		orderPaymentUseCase: paymentUseCase,
	}
}

func (o *OrderUseCase) OpenOrder(order dto.CustomerOrderRequest) (*entity.CustomerOrder, error) {
	var orderDetails []entity.CustomerOrderDetail
	for _, od := range order.Orders {
		orderDetail := entity.CustomerOrderDetail{
			MenuID: od.MenuId,
			Qty:    od.Qty,
			Price:  od.Price,
		}
		orderDetails = append(orderDetails, orderDetail)
	}
	ord, err := o.orderRepo.CreateOne(entity.CustomerOrder{
		CustomerName: order.CustomerName,
		OrderDetails: orderDetails,
		Model:        gorm.Model{},
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

func (o *OrderUseCase) CloseOrder(closeOrderInfo dto.CloseOrderRequest) (string, error) {
	total, err := o.orderRepo.GetSummaryPrice(closeOrderInfo.BillNo)
	log.Println(total)
	if err != nil {
		return "", err
	}
	closeOrderInfo.Total = total
	billId, err := o.orderPaymentUseCase.Payment(closeOrderInfo)
	if err != nil {
		return "", err
	}
	err = o.resvUseCase.CloseTable(billId.CustomerOrderID)
	if err != nil {
		return billId.CustomerOrderID, err
	}
	return billId.CustomerOrderID, nil
}
