package domain

import (
	"context"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

//Mark : represent every mark of a Test done by Users
type Mark struct {
	MarkID        int64          `json:"markID" gorm:"column:markID"`
	UserID        int64          `json:"userID" gorm:"column:userID"`
	QuestionaryID int64          `json:"questionaryID" gorm:"column:questionaryID"`
	Val           int            `json:"val" gorm:"column:val"`
	CreatedAt     mysql.NullTime `json:"created" gorm:"column:created"`
	UpdatedAt     mysql.NullTime `json:"updated" gorm:"column:updated"`
	DeletedAt     mysql.NullTime `json:"deleted" gorm:"column:deleted"` //Soft delete feature
}

//TableName : table name for Gorm
func (m *Mark) TableName() string {
	return "mark"
}

// MarkClient : defines the interface to access Test Marks data
type MarkClient interface {
	GetMark(ctx context.Context, userID int64, quetionaryID int64) (mark Mark, err error)
	GetMarkByLesson(ctx context.Context, userID int64, lessonID int64) (mark Mark, err error)
	GetMarkByQuestionary(ctx context.Context, userID int64, questionaryID int64) (mark Mark, err error)
	GetAllMarksByUser(ctx context.Context, userID int64) (marks []Mark, err error)
	GetAllMarks(ctx context.Context, courseID int64, language string) (marks []Mark, err error)
	GetAllMarksByLesson(ctx context.Context, lessonID Lesson) (marks []Mark, err error)
	CreateMark(ctx context.Context, mark *Mark) error
	UpdateMark(ctx context.Context, mark *Mark) error
	DeleteMark(ctx context.Context, userID int64, questionaryID int64) error
}

// asserts Client implements the MarkClient interface
var _ MarkClient = (*Client)(nil)

// GetMark :
func (c *Client) GetMark(ctx context.Context, userID int64, quetionaryID int64) (mark Mark, err error) {
	mk := Mark{}
	err = c.db.Table("mark").Where("userID = ? AND questionaryID = ?", userID, quetionaryID).Find(&mk).Error
	if err != nil {
		fmt.Println(err)
	}
	return mk, nil
}

// GetMarkByLesson :
func (c *Client) GetMarkByLesson(ctx context.Context, userID int64, lessonID int64) (mark Mark, err error) {
	q, err := c.GetQuestionaryByLessonID(ctx, lessonID)
	if err != nil {
		fmt.Println(err)
	}

	mk := Mark{}

	err = c.db.Table("mark").Where("userID = ? AND questionaryID = ?", userID, q.QuestionaryID).Find(&mk).Error
	if err != nil {
		fmt.Println(err)
	}
	return mk, nil
}

// GetMarkByQuestionary :
func (c *Client) GetMarkByQuestionary(ctx context.Context, userID int64, questionaryID int64) (mark Mark, err error) {
	mk := Mark{}

	err = c.db.Table("mark").Where("userID = ? AND questionaryID = ?", userID, questionaryID).Find(&mk).Error
	if err != nil {
		fmt.Println(err)
	}
	return mk, nil
}

// GetAllMarksByUser :
func (c *Client) GetAllMarksByUser(ctx context.Context, userID int64) (marks []Mark, err error) {
	mrks := []Mark{}
	err = c.db.Table("mark").Where("userID = ?", userID).Find(&mrks).Error
	if err != nil {
		fmt.Println(err)
	}
	return mrks, nil
}

// GetAllMarks :
func (c *Client) GetAllMarks(ctx context.Context, courseID int64, language string) (marks []Mark, err error) {
	//TODO
	return
}

// GetAllMarksByLesson :
func (c *Client) GetAllMarksByLesson(ctx context.Context, lessonID Lesson) (marks []Mark, err error) {
	mrks := []Mark{}
	err = c.db.Table("marks").Where("lessonID = ?", lessonID).Find(&mrks).Error
	if err != nil {
		fmt.Println(err)
		return mrks, nil
	}
	return
}

// CreateMark :
func (c *Client) CreateMark(ctx context.Context, mark *Mark) error {
	err := c.db.Create(&mark).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateMark :
func (c *Client) UpdateMark(ctx context.Context, mark *Mark) error {
	err := c.db.Save(&mark).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteMark :
func (c *Client) DeleteMark(ctx context.Context, userID int64, questionaryID int64) error {
	if userID == 0 || questionaryID == 0 {
		return fmt.Errorf("Error!!! (DeleteMark), incorrect UserID ID: %d | QuestionaryID: %d", userID, questionaryID)
	}

	mark := Mark{
		UserID:        userID,
		QuestionaryID: questionaryID,
	}
	return c.db.Delete(&mark).Error
}
