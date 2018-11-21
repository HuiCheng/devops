package configuration

import (
	"github.com/HuiCheng/devops/backend/util"
	"github.com/jinzhu/gorm"
)

// Config Model
type Config struct {
	ID         uint           `json:"id" gorm:"primary_key"`
	CreatedAt  util.JSONTime  `json:"-"`
	UpdatedAt  util.JSONTime  `json:"-"`
	DeletedAt  *util.JSONTime `json:"-" sql:"index" `
	Lable      string         `json:"lable"`
	Note       string         `json:"note"`
	Namespaces []Namespace    `json:"namespaces"`
	Keys       []Key          `json:"keys"`
	Values     []Value        `json:"values"`
}

// Namespace Model
type Namespace struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	CreatedAt util.JSONTime  `json:"-"`
	UpdatedAt util.JSONTime  `json:"-"`
	DeletedAt *util.JSONTime `json:"-" sql:"index" `
	Lable     string         `json:"lable" binding:"required,min=3,max=50"`
	ConfigID  uint           `json:"-"`
}

// Key Model
type Key struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	CreatedAt util.JSONTime  `json:"-"`
	UpdatedAt util.JSONTime  `json:"-"`
	DeletedAt *util.JSONTime `json:"-" sql:"index" `
	Lable     string         `json:"lable" binding:"required,min=3,max=50"`
	ConfigID  uint           `json:"-"`
}

// Value Model
type Value struct {
	gorm.Model
	Namespace   Namespace `json:"namespace"`
	NamespaceID uint      `json:"-"`
	Key         Key       `json:"key"`
	KeyID       uint      `json:"-"`
	Value       string    `json:"value"`
	ConfigID    uint      `json:"-"`
	GroupID     uint      `json:"-"`
}

// Group Model
type Group struct {
	gorm.Model
	Lable  string  `json:"lable"`
	Values []Value `json:"values"`
}
