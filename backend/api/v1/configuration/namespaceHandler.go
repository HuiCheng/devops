package configuration

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetNamespace Get ALL
func GetNamespace(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"data": "",
		},
	)
}

// PostNamespace New One
func PostNamespace(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusCreated,
			"data": "",
		},
	)
}
