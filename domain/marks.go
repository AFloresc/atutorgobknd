package domain

import (
	"context"

	"github.com/go-sql-driver/mysql"
)

//Mark : represent every mark of a Test done by Users
type Mark struct {
	UserID    int64          `json:"userID" gorm:"column:userID;primary_key"`
	TestID    int64          `json:"testID" gorm:"column:testID;primary_key"`
	value     int            `json:"value" gorm:"column:value"`
	CreatedAt mysql.NullTime `json:"created" gorm:"column:created"`
	UpdatedAt mysql.NullTime `json:"updated" gorm:"column:updated"`
	DeletedAt mysql.NullTime `json:"deleted" gorm:"column:deleted"` //Soft delete feature
}

//TableName : table name for Gorm
func (m *Mark) TableName() string {
	return "marks"
}

// MarkClient : defines the interface to access Test Marks data
type MarkClient interface {
	GetAllMarksByUser(ctx context.Context, userID int64) (marks []Mark, err error)
	GetAllMarks(ctx context.Context, courseID int64, language string) (marks []Mark, err error)
	GetAllMarksByLesson(ctx context.Context, lessonID Lesson) error
	CreateMark(ctx context.Context, mark *Mark) error
	UpdateMark(ctx context.Context, mark *Mark) error
	DeleteMark(ctx context.Context, mark *Mark) error
}

// asserts Client implements the MarkClient interface
var _ MarkClient = (*Client)(nil)

func (c *Client) GetAllMarksByUser(ctx context.Context, userID int64) (marks []Mark, err error) {
	//TODO
	return
}
func (c *Client) GetAllMarks(ctx context.Context, courseID int64, language string) (marks []Mark, err error) {
	//TODO
	return
}

func (c *Client) GetAllMarksByLesson(ctx context.Context, lessonID Lesson) error {
	//TODO
	return nil
}

func (c *Client) CreateMark(ctx context.Context, mark *Mark) error {
	//TODO
	return nil
}

func (c *Client) UpdateMark(ctx context.Context, mark *Mark) error {
	//TODO
	return nil
}

func (c *Client) DeleteMark(ctx context.Context, mark *Mark) error {
	//TODO
	return nil
}
