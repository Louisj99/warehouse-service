package usecases

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type getAllItemsResponse struct {
	BarcodePrefix string `json:"barcodePrefix"`
	ItemName      string `json:"itemName"`
	Description   string `json:"description"`
	LocationID    string `json:"locationID"`
	Quantity      int    `json:"quantity"`
}

func GetAllItems(warehouseRepository WarehouseRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		dbItemsReturn, err := warehouseRepository.GetAllItems(c)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		var response []getAllItemsResponse
		for _, dbItemReturn := range dbItemsReturn {
			response = append(response, getAllItemsResponse{
				BarcodePrefix: dbItemReturn.BarcodePrefix,
				ItemName:      dbItemReturn.ItemName,
				Description:   dbItemReturn.Description,
				LocationID:    dbItemReturn.LocationName,
				Quantity:      dbItemReturn.Quantity,
			})
		}
		c.JSON(http.StatusOK, response)
	}
}
