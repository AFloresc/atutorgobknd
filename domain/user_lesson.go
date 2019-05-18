package domain

import (
	"context"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

//Lesson represent every box inside the app
type UserLessons struct {
	UserID    int64          `json:"userID" gorm:"column:userID"`
	LessonID  int64          `json:"lessonID" gorm:"column:lessonID"`
	CreatedAt mysql.NullTime `json:"created" gorm:"column:created"`
	UpdatedAt mysql.NullTime `json:"updated" gorm:"column:updated"`
	DeletedAt mysql.NullTime `json:"deleted" gorm:"column:deleted"` //Soft delete feature
}

//TableName : table name for Gorm
func (ul *UserLessons) TableName() string {
	return "user_lessons"
}

// UserLessonsClient : defines the interface to access UserLessons data
type UserLessonsClient interface {
	GetUserLessons(ctx context.Context, userID int64, lessonID int64) (userlessons UserLessons, err error)
	GetAllUserLessonsByCourseID(ctx context.Context, courseID int64, userID int64, lessonID int64) (lesson []Lesson, err error)
	CreateUserLessons(ctx context.Context, userlessons *UserLessons) error
	UpdateUserLessons(ctx context.Context, userlessons *UserLessons) error
	DeleteUserLessons(ctx context.Context, userID int64, lessonID int64) error
}

// asserts Client implements the LessonClient interface
var _ LessonClient = (*Client)(nil)

// GetUserLessons : retrieves a userlesson row by it's userID and lessonID
func (c Client) GetUserLessons(ctx context.Context, userID int64, lessonID int64) (userlessons UserLessons, err error) {
	uLesson := UserLessons{}
	err = c.db.Table("user_lessons").Where("userID = ? AND lessonID = ?", userID, lessonID).Find(&uLesson).Error
	if err != nil {
		fmt.Println(err)
	}
	return uLesson, nil
}

func (c Client) GetAllUserLessonsByCourseID(ctx context.Context, courseID int64, userID int64, lessonID int64) (lesson []Lesson, err error) {
	//TODO
	return
}

func (c Client) CreateUserLessons(ctx context.Context, userlessons *UserLessons) error {
	//TODO
	return nil
}
func (c Client) UpdateUserLessons(ctx context.Context, userlessons *UserLessons) error {
	err := c.db.Save(userlessons).Error
	if err != nil {
		return err
	}
	return nil
}
func (c Client) DeleteUserLessons(ctx context.Context, userID int64, lessonID int64) error {
	if lessonID == 0 {
		return fmt.Errorf("Error!!! (DeleteLesson), incorrect Lesson ID: %d", lessonID)
	}

	lesson := Lesson{
		LessonID: lessonID,
	}
	return c.db.Delete(&lesson).Error

}
