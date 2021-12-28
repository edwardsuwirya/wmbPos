package usecase

import (
	"github.com/edwardsuwirya/wmbPos/dto"
	"github.com/edwardsuwirya/wmbPos/entity"
	"github.com/edwardsuwirya/wmbPos/repository"
)

type IOrderPaymentUseCase interface {
	Payment(payInfo dto.CloseOrderRequest) (*entity.OrderPayment, error)
}

type OrderPaymentUseCase struct {
	paymentRepo repository.IOrderPaymentRepository
	opoPayment  repository.IOpoPaymentRepository
}

func NewOrderPaymentUseCase(paymentRepo repository.IOrderPaymentRepository, opoRepo repository.IOpoPaymentRepository) IOrderPaymentUseCase {
	return &OrderPaymentUseCase{
		paymentRepo: paymentRepo,
		opoPayment:  opoRepo,
	}
}

func (r *OrderPaymentUseCase) Payment(payInfo dto.CloseOrderRequest) (*entity.OrderPayment, error) {
	paymentMethod, err := r.paymentRepo.GetPaymentMethodById(payInfo.PaymentMethod)
	if err != nil {
		return nil, err
	}
	if paymentMethod.PaymentMethodName == "OPO" {
		receipt, err := r.opoPayment.Payment(payInfo.PhoneNo, payInfo.Total)
		if err != nil {
			return nil, err
		}
		return r.paymentRepo.CreateOne(entity.OrderPayment{
			CustomerOrderID: payInfo.BillNo,
			BillerReceipt:   receipt,
			PaymentID:       payInfo.PaymentMethod,
		})
	}

	return r.paymentRepo.CreateOne(entity.OrderPayment{
		CustomerOrderID: payInfo.BillNo,
		BillerReceipt:   "",
		PaymentID:       payInfo.PaymentMethod,
	})
}
