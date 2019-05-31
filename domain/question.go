package domain

import (
	"context"
	"fmt"

	"github.com/atutor/utils"
	"github.com/go-sql-driver/mysql"
)

//Question represents every question inside questionary
type Question struct {
	QuestionID    int64          `json:"questionID" gorm:"column:questionID;primary_key"`
	QuestionaryID int64          `json:"questionaryID" gorm:"column:questionaryID"`
	Description   string         `json:"description" gorm:"column:description"`
	CreatedAt     mysql.NullTime `json:"created" gorm:"column:created"`
	UpdatedAt     mysql.NullTime `json:"updated" gorm:"column:updated"`
	DeletedAt     mysql.NullTime `json:"deleted" gorm:"column:deleted"` //Soft delete feature
}

//TableName : table name for Gorm
func (t *Question) TableName() string {
	return "question"
}

// QuestionClient : defines the interface to access Quertion data
type QuestionClient interface {
	GetQuestionsByQuestionaryID(ctx context.Context, questionaryID int64) (questions []Question, err error)
	CreateQuestion(ctx context.Context, question *Question) error
	UpdateQuestion(ctx context.Context, question *Question) error
	DeleteQuestion(ctx context.Context, testID int64) error
}

// asserts Client implements the TestClient interface
var _ QuestionClient = (*Client)(nil)

func (c Client) GetQuestionsByQuestionaryID(ctx context.Context, questionaryID int64) (questions []Question, err error) {
	q := []Question{}
	err = c.db.Table("question").Where("QuestionaryID = ?", questionaryID).Find(&q).Error
	if err != nil {
		fmt.Println(err)
	}
	return q, nil
}

// CreateQuestion :
func (c Client) CreateQuestion(ctx context.Context, question *Question) error {
	err := c.db.Create(&question).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateQuestion :
func (c Client) UpdateQuestion(ctx context.Context, question *Question) error {
	err := c.db.Save(&question).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteQuestion :
func (c Client) DeleteQuestion(ctx context.Context, questionID int64) error {
	if questionID == 0 {
		return fmt.Errorf("Error!!! (DeleteQuestion), incorrect Question ID: %d", questionID)
	}

	question := Question{
		QuestionID: questionID,
	}
	return c.db.Delete(&question).Error
}

// hardDeleteQuestion :
func (c Client) hardDeleteQuestion(ctx context.Context, questionID int64) error {
	if questionID != 0 {
		return c.db.Exec("DELETE FROM question WHERE questionID=? ", questionID).Error
	}
	return utils.NewError("Question ID value not allowed on hard Delete action")
}
