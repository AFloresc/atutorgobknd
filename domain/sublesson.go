package domain

import (
	"github.com/go-sql-driver/mysql"
)

//Sublesson:
type Sublesson struct {
	SublessonID int64          `json:"sublessonID" gorm:"column:id;primary_key"`
	LessonName  string         `json:"lessonName" gorm:"column:lessonName"`
	Content     string         `json:"content" gorm:"column:content;type:text"`
	Image       string         `json:"image" gorm:"column:Image;type:text"`
	CreatedAt   mysql.NullTime `json:"created" gorm:"column:created"`
	UpdatedAt   mysql.NullTime `json:"updated" gorm:"column:updated"`
}
