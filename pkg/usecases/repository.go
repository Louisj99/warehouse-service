package usecases

import (
	"context"
	"warehouse-service/pkg/entities"
)

//go:generate mockgen --build_flags=--mod=mod -destination=../../mocks/usecases/warehouseRepository.go  .
type WarehouseRepository interface {
	PlaceholderAdapter(ctx context.Context, placeholder string) error
	GetItemInformation(ctx context.Context, barcode string) (*entities.ItemInformation, error)
	GetAllItems(ctx context.Context) ([]*entities.ItemInformation, error)
	UpdateItemInformation(ctx context.Context, barcode string, name string, description string, location string, quantity int) error
}
