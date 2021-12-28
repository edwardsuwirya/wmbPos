package manager

import "github.com/edwardsuwirya/wmbPos/usecase"

type UseCaseManager interface {
	OrderUseCase() usecase.IOrderUseCase
	TableReservationUseCase() usecase.IOrderTableReservationUseCase
	PaymentUseCase() usecase.IOrderPaymentUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (uc *useCaseManager) OrderUseCase() usecase.IOrderUseCase {
	return usecase.NewOrderUseCase(uc.repo.OrderRepo(), uc.TableReservationUseCase(), uc.PaymentUseCase())
}
func (uc *useCaseManager) TableReservationUseCase() usecase.IOrderTableReservationUseCase {
	return usecase.NewOrderTableReservationUseCase(uc.repo.ResvRepo())
}
func (uc *useCaseManager) PaymentUseCase() usecase.IOrderPaymentUseCase {
	return usecase.NewOrderPaymentUseCase(uc.repo.PaymentRepo(), uc.repo.OpoPaymentRepo())
}
func NewUseCaseManger(manager RepoManager) UseCaseManager {
	return &useCaseManager{repo: manager}
}
