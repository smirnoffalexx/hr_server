package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SwaggerHandler returns the Swagger handler
func SwaggerHandler() gin.HandlerFunc {
	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}
