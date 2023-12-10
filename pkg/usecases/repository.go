package usecases

import "context"

type WarehouseRepository interface {
	PlaceholderAdapter(ctx context.Context, placeholder string) error
}
