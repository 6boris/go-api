package service

import (
	"github.com/dicf/go-api/model/mysql"
	"github.com/jinzhu/gorm"
	"strconv"
)

func GetTag(id int) *mysql.Tag {
	db := db_mysql

	tag := mysql.Tag{}
	db.Where("id = ?", id).Find(&tag)

	return &tag
}

func GetTags(filter map[string]interface{}) []mysql.Tag {
	db := filterTags(filter)
	tags := []mysql.Tag{}

	db.Find(&tags)

	return tags
}

func CreatedTag(filter map[string]interface{}) *mysql.Tag {
	db := db_mysql

	if _, ok := filter["name"]; !ok && filter["name"] == "" {
		return nil
	}

	tag := mysql.Tag{}
	tag.Name = filter["name"].(string)
	db.Save(&tag)

	return &tag
}

func UpdateTag(filter map[string]string) *mysql.Tag {
	db := db_mysql
	id, _ := strconv.Atoi(filter["id"])
	tag := mysql.Tag{}
	tag.Id = id
	tag.Name = filter["name"]

	db.Save(&tag)
	return &tag
}

func DeleteTag(id int) int {
	db := db_mysql
	tag := mysql.Tag{}
	tag.Id = id

	db.Delete(&tag)

	return tag.Id
}

func filterTags(filter map[string]interface{}) *gorm.DB {
	db := db_mysql

	if _, ok := filter["name"]; ok && filter["name"] != "" {
		db = db.Where("name like ?", "%"+filter["name"].(string)+"%")
	}
	return db

}
