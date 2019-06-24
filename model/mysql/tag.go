package mysql

import (
	"github.com/dicf/go-api/pkg/json"
)

type Tag struct {
	Id        int            `gorm:"column:id" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`
	CreatedAt *json.JsonTime `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *json.JsonTime `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *json.JsonTime `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedBy int            `gorm:"column:created_by" json:"created_by"`
	UpdatedBy int            `gorm:"column:updated_by" json:"updated_by"`
	Status    int            `gorm:"column:status" json:"status"`
}

func (self Tag) TableName() string {
	return "blog_tag"
}
