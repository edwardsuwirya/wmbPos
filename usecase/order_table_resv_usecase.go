package usecase

import (
	"github.com/edwardsuwirya/wmbPos/dto"
	"github.com/edwardsuwirya/wmbPos/repository"
)

type IOrderTableReservationUseCase interface {
	ReserveTable(req dto.TableRequest) error
	CloseTable(billNo string) error
}

type OrderTableReservationUseCase struct {
	resvRepo repository.ITableOrderReservationRepository
}

func NewOrderTableReservationUseCase(resvRepo repository.ITableOrderReservationRepository) IOrderTableReservationUseCase {
	return &OrderTableReservationUseCase{
		resvRepo: resvRepo,
	}
}

func (r *OrderTableReservationUseCase) ReserveTable(req dto.TableRequest) error {
	return r.resvRepo.ReserveOne(req)
}

func (r *OrderTableReservationUseCase) CloseTable(billNo string) error {
	return r.resvRepo.Close(billNo)
}
