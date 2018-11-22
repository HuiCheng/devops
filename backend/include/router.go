package include

import (
	"net/http"

	"github.com/HuiCheng/devops/backend/api/v1/configuration"
	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

// Database Middleware
func Database(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}

// MKRouter Func
func MKRouter() *gin.Engine {
	router := gin.Default()
	router.Use(Database(DB))

	router.GET("/HealthCheck",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		},
	)

	v1 := router.Group("/v1")
	v1C := v1.Group("/configuration")
	v1CN := v1C.Group("/namespace")
	{
		v1CN.GET("", configuration.GetNamespaceHandler)
		v1CN.POST("", configuration.PostNamespaceHandler)
	}
	v1CK := v1C.Group("/key")
	{
		v1CK.POST("", configuration.PostKeyHandler)
	}

	v1CC := v1C.Group("/config")
	{
		v1CC.GET("", configuration.GetConfigHandler)
		v1CC.POST("", configuration.PostConfigHandler)
	}

	return router
}
