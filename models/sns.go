package models

import (
	"gorm.io/gorm"
)

type Sns struct {
	gorm.Model
	Name string
	//	Enable bool
}
