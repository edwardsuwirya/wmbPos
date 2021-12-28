package manager

import "github.com/edwardsuwirya/wmbPos/repository"

type RepoManager interface {
	OrderRepo() repository.IOrderRepository
	ResvRepo() repository.ITableOrderReservationRepository
	PaymentRepo() repository.IOrderPaymentRepository
	OpoPaymentRepo() repository.IOpoPaymentRepository
}

type repoManager struct {
	infra Infra
}

func (rm *repoManager) OrderRepo() repository.IOrderRepository {
	return repository.NewOrderRepository(rm.infra.SqlDb())
}

func (rm *repoManager) ResvRepo() repository.ITableOrderReservationRepository {

	return repository.NewTableOrderReservation(rm.infra.HttpClient(), rm.infra.Config().TableManagementConfig)
}

func (rm *repoManager) PaymentRepo() repository.IOrderPaymentRepository {
	return repository.NewOrderPaymentRepository(rm.infra.SqlDb())
}

func (rm *repoManager) OpoPaymentRepo() repository.IOpoPaymentRepository {
	return repository.NewOpoPaymentRepository(rm.infra.HttpClient(), rm.infra.Config().OpoPaymentConfig)
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{infra}
}
