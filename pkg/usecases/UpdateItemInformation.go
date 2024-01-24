package usecases

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type updateSingleItemInformationRequest struct {
	BarcodePrefix string `json:"barcodePrefix" required:"true"`
	ItemName      string `json:"itemName" required:"true"`
	Description   string `json:"description" required:"true"`
	LocationName  string `json:"locationName" required:"true"`
	Quantity      int    `json:"quantity" required:"true"`
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
		err = warehouseRepository.UpdateItemInformation(c, request.BarcodePrefix, request.ItemName, request.Description, request.LocationName, request.Quantity)
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
