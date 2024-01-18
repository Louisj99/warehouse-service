package adapters

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"warehouse-service/pkg/entities"
)

type PostgresAdapter struct {
	DB *sql.DB
}

func NewPostgresAdapter(connStr string) (*PostgresAdapter, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresAdapter{DB: db}, nil
}

func (p *PostgresAdapter) PlaceholderAdapter(ctx context.Context, placeholder string) error {
	//TODO implement me
	panic("implement me")
}

func (p *PostgresAdapter) GetItemInformation(ctx context.Context, barcode int) (*entities.ItemInformation, error) {
	query := `SELECT Item.item_id, Item.item_name, Item.description, Inventory.location_id 
              FROM Item 
              INNER JOIN Inventory ON Item.item_id = Inventory.item_id 
              WHERE Item.item_id = $1`

	row := p.DB.QueryRowContext(ctx, query, barcode)

	item := &entities.ItemInformation{}
	err := row.Scan(&item.BarcodePrefix, &item.ItemName, &item.Description, &item.LocationID)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, nil
		}
		return nil, err
	}
	return item, nil
}
