package usecases

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type getSingleItemInformationRequest struct {
	Barcode string `json:"barcode" required:"true"`
}

type getSingleItemInformationResponse struct {
	BarcodePrefix string `json:"barcodePrefix"`
	ItemName      string `json:"itemName"`
	Description   string `json:"description"`
	LocationName  string `json:"locationName"`
	Quantity      int    `json:"quantity"`
}

func GetSingleItemInformation(warehouseRepository WarehouseRepository) gin.HandlerFunc {
	return func(c *gin.Context) {

		var request getSingleItemInformationRequest

		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(400, gin.H{"Error With Request": err.Error()})
			return
		}
		if request.Barcode == "" {
			c.JSON(400, gin.H{"error": "barcode is required"})
			return
		}
		dbItemReturn, err := warehouseRepository.GetItemInformation(c, request.Barcode)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				c.JSON(404, gin.H{"error": "item not found"})
				return
			}
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		response := getSingleItemInformationResponse{
			BarcodePrefix: dbItemReturn.BarcodePrefix,
			ItemName:      dbItemReturn.ItemName,
			Description:   dbItemReturn.Description,
			LocationName:  dbItemReturn.LocationName,
			Quantity:      dbItemReturn.Quantity,
		}
		c.JSON(http.StatusOK, response)
	}
}
