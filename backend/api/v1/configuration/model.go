package configuration

import (
	"time"
)

// Config Model
type Config struct {
	ID         uint        `json:"id" gorm:"primary_key"`
	CreatedAt  time.Time   `json:"-" time_format:"2006-01-02 15:04:05"`
	UpdatedAt  time.Time   `json:"-"`
	DeletedAt  *time.Time  `json:"-" sql:"index"`
	Lable      string      `json:"lable"`
	Note       string      `json:"note"`
	Namespaces []Namespace `json:"namespaces"`
	Keys       []Key       `json:"keys"`
	Values     []Value     `json:"values"`
}

// Namespace Model
type Namespace struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Lable     string     `json:"lable" binding:"required,min=3,max=50"`
	ConfigID  uint       `json:"-"`
	ValueID   uint       `json:"-"`
}

// Key Model
type Key struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Lable     string     `json:"lable" binding:"required,min=3,max=50"`
	ConfigID  uint       `json:"-"`
	ValueID   uint       `json:"-"`
}

// Value Model
type Value struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	Namespace Namespace  `json:"namespace"`
	Key       Key        `json:"key"`
	Value     string     `json:"value"`
	ConfigID  uint       `json:"-"`
}
