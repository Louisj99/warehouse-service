package adapters

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"time"
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

func (p *PostgresAdapter) GetItemInformation(ctx context.Context, barcode string) (*entities.ItemInformation, error) {
	query := `SELECT Item.item_id, Item.item_name, Item.description, Inventory.quantity, Location.location_name
			  FROM Item 
			  INNER JOIN Inventory ON Item.item_id = Inventory.item_id
			  INNER JOIN Location ON Inventory.location_id = Location.location_id
              WHERE Item.item_id = $1`

	row := p.DB.QueryRowContext(ctx, query, barcode)

	item := &entities.ItemInformation{}
	err := row.Scan(&item.BarcodePrefix, &item.ItemName, &item.Description, &item.Quantity, &item.LocationName)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (p *PostgresAdapter) GetAllItems(ctx context.Context) ([]entities.ItemInformation, error) {
	query := `SELECT Item.item_id, Item.item_name, Item.description, Inventory.quantity, Location.location_name
			  FROM Item 
			  INNER JOIN Inventory ON Item.item_id = Inventory.item_id
			  INNER JOIN Location ON Inventory.location_id = Location.location_id`

	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var items []entities.ItemInformation
	for rows.Next() {
		item := entities.ItemInformation{}
		err := rows.Scan(&item.BarcodePrefix, &item.ItemName, &item.Description, &item.Quantity, &item.LocationName)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (p *PostgresAdapter) UpdateItemInformation(ctx context.Context, barcodePrefix string, itemName string, description string, locationName string, quantity int, updatedAt time.Time) error {
	var locationID int

	getLocationIDQuery := `SELECT location_id FROM Location WHERE location_name = $1;`
	row := p.DB.QueryRowContext(ctx, getLocationIDQuery, locationName)
	err := row.Scan(&locationID)
	if err != nil {
		return err
	}

	updateInventoryQuery := `UPDATE Inventory
SET location_id = $1, quantity = $2, last_updated = $3
WHERE item_id = $4;`
	updateItemQuery := `UPDATE item
	SET item_name = $1, description = $2
	WHERE item_id = $3;`

	_, err = p.DB.Exec(updateInventoryQuery, locationID, quantity, updatedAt, barcodePrefix)
	if err != nil {
		return err
	}
	_, err = p.DB.Exec(updateItemQuery, itemName, description, barcodePrefix)

	return nil
}
