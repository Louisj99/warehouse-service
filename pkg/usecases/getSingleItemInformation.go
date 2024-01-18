package usecases

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type getSingleItemInformationRequest struct {
	barcode int `json:"barcode"`
}

type getSingleItemInformationResponse struct {
	barcodePrefix int    `json:"barcodePrefix"`
	itemName      string `json:"itemName"`
	description   string `json:"description"`
	locationID    int    `json:"locationID"`
	quantity      int    `json:"quantity"`
}

func GetSingleItemInformation(warehouseRepository WarehouseRepository) gin.HandlerFunc {
	return func(c *gin.Context) {

		var request getSingleItemInformationRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		dbItemReturn, err := warehouseRepository.GetItemInformation(c, request.barcode)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		response := getSingleItemInformationResponse{
			barcodePrefix: dbItemReturn.BarcodePrefix,
			itemName:      dbItemReturn.ItemName,
			description:   dbItemReturn.Description,
			locationID:    dbItemReturn.LocationID,
		}
		c.JSON(http.StatusOK, response)
	}
}
