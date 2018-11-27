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
	if err := db.Preload("Namespaces").Preload("Keys").Preload("Values").
		Preload("Values.Namespace").Preload("Values.Key").
		Find(&result).Error; err != nil {
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

	for i := range post.Values {
		err1 := db.Model(&post.Values[i].Namespace).Update("ValueID", post.Values[i].ID).Error
		err2 := db.Model(&post.Values[i].Key).Update("ValueID", post.Values[i].ID).Error
		if err1 != nil || err2 != nil {
			glog.Errorln(err1.Error(), err2.Error())
			c.JSON(
				http.StatusBadRequest,
				gin.H{"code": http.StatusBadRequest, "data": []string{err1.Error(), err2.Error()}},
			)
		}
	}

	c.JSON(
		http.StatusOK,
		gin.H{"code": http.StatusCreated, "data": ""},
	)

}
