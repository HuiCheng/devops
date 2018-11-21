package configuration

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// PostKeyHandler New One
func PostKeyHandler(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	var post Key
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "data": err.Error()},
		)
	} else if err := db.Create(&post).Error; err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "data": err.Error()},
		)
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"code": http.StatusCreated, "data": post},
		)
	}
}
