package usecases

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type updateSingleItemInformationRequest struct {
	BarcodePrefix string  `json:"barcodePrefix"`
	ItemName      string  `json:"itemName" `
	Description   string  `json:"description" `
	LocationName  string  `json:"locationName"`
	Quantity      int     `json:"quantity"`
	Price         float32 `json:"price"`
	Category      string  `json:"category"`
	Size          int     `json:"size" `
}

func UpdateItemInformation(warehouseRepository WarehouseRepository) gin.HandlerFunc {
	return func(c *gin.Context) {

		var request updateSingleItemInformationRequest

		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(400, gin.H{"Error With Request": err.Error()})
			return
		}
		if request.BarcodePrefix == "" {
			c.JSON(400, gin.H{"error": "barcodePrefix is required"})
			return
		}
		if request.ItemName == "" {
			c.JSON(400, gin.H{"error": "itemName is required"})
			return
		}
		if request.Description == "" {
			c.JSON(400, gin.H{"error": "description is required"})
			return
		}
		if request.LocationName == "" {
			c.JSON(400, gin.H{"error": "locationName is required"})
			return
		}
		if request.Quantity == 0 {
			c.JSON(400, gin.H{"error": "quantity is required"})
			return
		}
		if request.Price == 0 {
			c.JSON(400, gin.H{"error": "price is required"})
			return
		}
		if request.Category == "" {
			c.JSON(400, gin.H{"error": "category is required"})
			return
		}
		if request.Size == 0 {
			c.JSON(400, gin.H{"error": "size is required"})
			return
		}
		currentTime := time.Now()
		err = warehouseRepository.UpdateItemInformation(c, request.BarcodePrefix, request.ItemName, request.Description, request.Price, request.Category, request.Size, request.LocationName, request.Quantity, currentTime)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				c.JSON(404, gin.H{"error": "item not found"})
				return
			}
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "item updated"})
	}
}
