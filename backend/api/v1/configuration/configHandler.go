package configuration

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// PostConfigHandler New One
func PostConfigHandler(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)
	_ = db

	c.JSON(
		http.StatusOK,
		gin.H{"code": http.StatusCreated, "data": ""},
	)
}
