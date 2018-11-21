package configuration

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// NamespaceHandler Struct
type NamespaceHandler struct {
	DB *gorm.DB
}

// Get ALL
func (h *NamespaceHandler) Get(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusOK,
			"data": "",
		},
	)
}

// Post New One
func (h *NamespaceHandler) Post(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code": http.StatusCreated,
			"data": "",
		},
	)
}
