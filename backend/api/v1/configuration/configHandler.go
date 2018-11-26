package configuration

import (
	"fmt"
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
		return
	} else if err := db.Create(&post).Error; err != nil {
		glog.Errorln(err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{"code": http.StatusBadRequest, "data": err.Error()},
		)
		return
	}

	fmt.Println("Start")
	for i := range post.Values {
		if err := db.Model(&post.Values[i].Namespace).Update("ValueID", post.Values[i].ID).Error; err != nil {
			glog.Errorln(err.Error())
		}
		if err := db.Model(&post.Values[i].Key).Update("ValueID", post.Values[i].ID).Error; err != nil {
			glog.Errorln(err.Error())
		}
		fmt.Println(i)
	}
	fmt.Println("Send")

	c.JSON(
		http.StatusOK,
		gin.H{"code": http.StatusCreated, "data": ""},
	)

}
