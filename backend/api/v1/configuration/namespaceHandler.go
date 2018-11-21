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

	var post Namespace
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "data": err.Error()},
		)
		return
	}

	if err := db.Create(&post).Error; err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "data": err.Error()},
		)
		return
	}

	fmt.Println(post)
	c.JSON(
		http.StatusOK,
		gin.H{"code": http.StatusCreated, "data": post},
	)
}
