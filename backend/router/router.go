package router

import (
	"net/http"

	"github.com/HuiCheng/devops/backend/api/v1/configuration"
	"github.com/HuiCheng/devops/backend/include"

	"github.com/gin-gonic/gin"
)

// MKRouter Func
func MKRouter() *gin.Engine {
	router := gin.Default()
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
		namespace := configuration.NamespaceHandler{DB: include.DB}
		v1CN.GET("", namespace.Get)
		v1CN.POST("", namespace.Post)
	}

	return router
}
