package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetTotal(maps interface{}) (total int) {
	db.Model(&Tag{}).Where(maps).Count(&total)
	return
}

func ExistTagByName(maps interface{}) bool{
	var tag Tag
	db.Select("id").Where(maps).Find(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int, createBy string) bool {
	db.Create(&Tag{
		Name: name,
		State: state,
		CreatedBy: createBy,
	})
	return true
}

func ExistTagById(id int) bool {
	var tag Tag
	db.Select("id").Where(id).Find(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func EditTag(id int, data interface{}) {
	db.Model(&Tag{}).Where("id=?", id).Updates(data)
	return
}

func DeletedTag(id int) {
	db.Where("id=?", id).Delete(&Tag{})
	return
}

func (tag *Tag)BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("created_on", time.Now().Unix())
	return
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) {
	scope.SetColumn("modified_on", time.Now().Unix())
	return
}