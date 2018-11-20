package configuration

import (
	"github.com/jinzhu/gorm"
)

// Config Model
type Config struct {
	gorm.Model
	Lable      string      `json:"lable"`
	Namespaces []Namespace `json:"namespaces"`
	Keys       []Key       `json:"keys"`
	Values     []Value     `json:"values"`
}

// Namespace Model
type Namespace struct {
	gorm.Model
	Lable    string `json:"lable"`
	ConfigID uint   `json:"-"`
}

// Key Model
type Key struct {
	gorm.Model
	Lable    string `json:"lable"`
	ConfigID uint   `json:"-"`
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
