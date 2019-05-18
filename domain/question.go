package domain

import (
	"context"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

//Test represent every box inside the app
type Question struct {
	QuestionID  int64          `json:"questionID" gorm:"column:questionID;primary_key"`
	Description string         `json:"description" gorm:"column:description"`
	TestID      int64          `json:"testID" gorm:"column:test_id"`
	CreatedAt   mysql.NullTime `json:"created" gorm:"column:created"`
	UpdatedAt   mysql.NullTime `json:"updated" gorm:"column:updated"`
	DeletedAt   mysql.NullTime `json:"deleted" gorm:"column:deleted"` //Soft delete feature
}

//TableName : table name for Gorm
func (t *Question) TableName() string {
	return "tests"
}

// TestClient : defines the interface to access Test data
type QuestionClient interface {
	GetQuestionsByTestID(ctx context.Context, testID int64) (questions []Question, err error)
	CreateQuestion(ctx context.Context, question *Question) error
	UpdateQuestion(ctx context.Context, question *Question) error
	DeleteQuestion(ctx context.Context, testID int64) error
}

// asserts Client implements the TestClient interface
var _ QuestionClient = (*Client)(nil)

func (c Client) GetQuestionsByTestID(ctx context.Context, testID int64) (questions []Question, err error) {
	//TODO
	return nil, nil
}

func (c Client) CreateQuestion(ctx context.Context, question *Question) error {
	//TODO
	return nil
}
func (c Client) UpdateQuestion(ctx context.Context, question *Question) error {
	//TODO
	return nil
}
func (c Client) DeleteQuestion(ctx context.Context, questionID int64) error {
	if questionID == 0 {
		return fmt.Errorf("Error!!! (DeleteQuestion), incorrect Question ID: %d", questionID)
	}

	question := Question{
		QuestionID: questionID,
	}
	return c.db.Delete(&question).Error
}
