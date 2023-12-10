package drivers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"warehouse-service/pkg/usecases"
)

const (
	serviceName = "warehouse-service"
)

// SetupRouter configure the gine router and returns a pointer to a gin engine which can be run.
func SetupRouter(warehouseRepository usecases.WarehouseRepository) *gin.Engine {

	r := NewDefaultRouter(serviceName)

	// Configure routes
	v1 := r.Group("//v1")
	{
		v1.GET("/placeholder", usecases.Placeholder(warehouseRepository, "placeholder"))
	}

	return r
}

func NewDefaultRouter(serviceNameVal string, additionalMiddleware ...gin.HandlerFunc) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	r.Use(additionalMiddleware...)

	// Declares a config variable, assigning its valid methods and headers
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PATCH", "PUT", "OPTIONS", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Forwarded-Authorization", "Strict-Transport-Security"}

	r.Use(cors.New(config))

	return r
}
