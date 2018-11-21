package configuration

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetNamespaceHandler ALL
func GetNamespaceHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"data": "",
		},
	)
}

// PostNamespaceHandler New One
func PostNamespaceHandler(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)
	fmt.Println(db)

	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusCreated,
			"data": "",
		},
	)
}
