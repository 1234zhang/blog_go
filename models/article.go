package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model
	TagID 			int `json:"tag_id" gorm:"index"`
	Tag				string `json:"tag"`

	Title			string `json:"title"`
	Desc   			string `json:"desc"`
	Content			string `json:"content"`
	CreatedBy		string `json:"created_by"`
	ModifiedBy		string `json:"modified_by"`
	State 			int    `json:"state"`
}


func (article *Article) BeforeCreate(scope *gorm.Scope) {
	_ = scope.SetColumn("CreateOn", time.Now().Unix())
	return
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) {
	_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	return
}