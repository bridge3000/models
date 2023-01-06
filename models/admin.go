package models

import (
	//	"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	AdminName string `json:"admin_name"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
	GroupId   int    `json:"group_id"`
	GroupName string `gorm:"-"`
}

func (this *Admin) IsSuper() bool {
	return this.GroupId == 1
}
