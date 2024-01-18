package usecases

import (
	"context"
	"warehouse-service/pkg/entities"
)

type WarehouseRepository interface {
	PlaceholderAdapter(ctx context.Context, placeholder string) error
	GetItemInformation(ctx context.Context, barcode int) (*entities.ItemInformation, error)
}
