package manager

import "github.com/edwardsuwirya/wmbPos/repository"

type RepoManager interface {
	OrderRepo() repository.IOrderRepository
	ResvRepo() repository.ITableOrderReservationRepository
}

type repoManager struct {
	infra Infra
}

func (rm *repoManager) OrderRepo() repository.IOrderRepository {
	return repository.NewOrderRepository(rm.infra.SqlDb())
}

func (rm *repoManager) ResvRepo() repository.ITableOrderReservationRepository {
	return repository.NewTableOrderReservation()
}

func NewRepoManager(infra Infra) RepoManager {
	return &repoManager{infra}
}
