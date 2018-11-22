package configuration

import (
	"net/http"

	"github.com/golang/glog"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetConfigHandler All
func GetConfigHandler(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	var result []Config
	if err := db.Preload("Namespaces").Preload("Keys").Preload("Values").Find(&result).Error; err != nil {
		glog.Errorln(err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "data": err.Error()},
		)
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"code": http.StatusCreated, "data": result},
		)
	}
}

// PostConfigHandler New One
func PostConfigHandler(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	var post Config
	if err := c.ShouldBindJSON(&post); err != nil {
		glog.Errorln(err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "data": err.Error()},
		)
	} else if err := db.Create(&post).Error; err != nil {
		glog.Errorln(err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "data": err.Error()},
		)
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{"code": http.StatusCreated, "data": ""},
		)
	}

}
