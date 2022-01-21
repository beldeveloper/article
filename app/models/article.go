package model

import (
	"time"
	"gorm.io/datatypes"
)

type Article struct {
	Id          int64 `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	DateCreated time.Time
	Tags        datatypes.JSON
}