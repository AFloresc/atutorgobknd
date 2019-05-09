package domain

import (
	"context"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

//Lesson represent every box inside the app
type Lesson struct {
	LessonID  int64          `json:"lessonID" gorm:"column:lessonID;primary_key"`
	Title     string         `json:"titlee" gorm:"column:Title"`
	Text      string         `json:"text" gorm:"column:Text;type:text"`
	Image     string         `json:"image" gorm:"column:Image"`
	Language  string         `json:"language" gorm:"column:Language"`
	Contents  []Content      `json:"contents"`
	Test      Test           `json:"test"`
	CreatedAt mysql.NullTime `json:"created" gorm:"column:created"`
	UpdatedAt mysql.NullTime `json:"updated" gorm:"column:updated"`
	DeletedAt mysql.NullTime `json:"deleted" gorm:"column:deleted"` //Soft delete feature
	Position  int64          `json:"position" gorm:"column:position"`
}

//TableName : table name for Gorm
func (l *Lesson) TableName() string {
	return "lessons"
}

// LessonClient : defines the interface to access Lesson data
type LessonClient interface {
	GetLesson(ctx context.Context, lessonID int64) (lesson *Lesson, err error)
	GetAllLessonsByCourseID(ctx context.Context, courseID int64, language string) (lesson []Lesson, err error)
	CreateLesson(ctx context.Context, lesson *Lesson) error
	UpdateLesson(ctx context.Context, lesson *Lesson) error
	DeleteLesson(ctx context.Context, lessonID int64) error
}

// asserts Client implements the LessonClient interface
var _ LessonClient = (*Client)(nil)

func (c Client) GetLesson(ctx context.Context, lessonID int64) (lesson *Lesson, err error) {
	//TODO
	return nil, nil
}
func (c Client) CreateLesson(ctx context.Context, lesson *Lesson) error {
	//TODO
	return nil
}
func (c Client) UpdateLesson(ctx context.Context, lesson *Lesson) error {
	err := c.db.Save(lesson).Error
	if err != nil {
		return err
	}
	return nil
}
func (c Client) DeleteLesson(ctx context.Context, lessonID int64) error {
	if lessonID == 0 {
		return fmt.Errorf("Error!!! (DeleteLesson), incorrect Lesson ID: %d", lessonID)
	}

	lesson := Lesson{
		LessonID: lessonID,
	}
	return c.db.Delete(&lesson).Error

}

func (c Client) GetAllLessonsByCourseID(ctx context.Context, courseID int64, language string) (lesson []Lesson, err error) {
	//TODO
	return nil, nil
}
