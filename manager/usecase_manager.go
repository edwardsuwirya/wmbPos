package manager

import "github.com/edwardsuwirya/wmbPos/usecase"

type UseCaseManager interface {
	OrderUseCase() usecase.IOrderUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (uc *useCaseManager) OrderUseCase() usecase.IOrderUseCase {
	return usecase.NewOrderUseCase(uc.repo.OrderRepo(), uc.TableReservationUseCase())
}
func (uc *useCaseManager) TableReservationUseCase() usecase.IOrderTableReservationUseCase {
	return usecase.NewOrderTableReservationUseCase(uc.repo.ResvRepo())
}
func NewUseCaseManger(manager RepoManager) UseCaseManager {
	return &useCaseManager{repo: manager}
}
