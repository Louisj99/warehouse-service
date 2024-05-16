package usecases

import (
	"context"
	"time"
	"warehouse-service/pkg/entities"
)

//go:generate mockgen --build_flags=--mod=mod -destination=../../mocks/usecases/warehouseRepository.go  . WarehouseRepository
type WarehouseRepository interface {
	PlaceholderAdapter(context.Context, string) error
	GetItemInformation(context.Context, string) (*entities.ItemInformation, error)
	GetAllItems(ctx context.Context) ([]entities.ItemInformation, error)
	UpdateItemInformation(context.Context, string, string, string, string, int, time.Time) error
}

//TODO : find a reason to not add this to the usecase files individually
