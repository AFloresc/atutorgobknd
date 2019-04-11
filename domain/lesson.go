package domain

import (
	"context"

	"github.com/go-sql-driver/mysql"
)

//Lesson represent every box inside the app
type Lesson struct {
	LessonID   int64          `json:"lessonID" gorm:"column:lessonID;primary_key"`
	Title      string         `json:"titlee" gorm:"column:Title"`
	Text       string         `json:"text" gorm:"column:Text;type:text"`
	Image      string         `json:"image" gorm:"column:Image"`
	UserID     int64          `json:"userID" gorm:"column:userID"`
	CourseID   int64          `json:"courseID" gorm:"column:courseID"`
	Sublessons []Sublesson    `json:"sublessons"`
	CreatedAt  mysql.NullTime `json:"created" gorm:"column:created"`
	UpdatedAt  mysql.NullTime `json:"updated" gorm:"column:updated"`
}

//TableName : table name for Gorm
func (Lesson) TableName() string {
	return "lessons"
}

// LessonClient : defines the interface to access Lesson data
type LessonClient interface {
	GetLesson(ctx context.Context, lessonID int64) (lesson *Lesson, err error)
	GetAllLessonsByCourseID(ctx context.Context, lessonID int64) (lesson []Lesson, err error)
	CreateLesson(ctx context.Context, lesson *Lesson) error
	UpdateLesson(ctx context.Context, lesson *Lesson) error
	DeleteLesson(ctx context.Context, lessonID int64) error
	GetAllSublessonsByCourse(ctx context.Context, courseid int64) ([]Lesson, error)
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
	//TODO
	return nil
}
func (c Client) DeleteLesson(ctx context.Context, lessonID int64) error {
	//TODO
	return nil
}
func (c Client) GetAllSublessonsByCourse(ctx context.Context, courseid int64) ([]Lesson, error) {
	//TODO
	return nil, nil
}

func (c Client) GetAllLessonsByCourseID(ctx context.Context, lessonID int64) (lesson []Lesson, err error) {
	//TODO
	return nil, nil
}
