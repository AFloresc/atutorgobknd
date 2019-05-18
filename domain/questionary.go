package domain

import (
	"context"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

//Questionary represent every box inside the app
type Questionary struct {
	QuestionaryID int64          `json:"questionaryID" gorm:"column:QuestionaryID;primary_key"`
	Description   string         `json:"description" gorm:"column:description"`
	LessonID      int64          `json:"QuestionaryID" gorm:"column:Questionary_id"`
	Questions     []Question     `json:"questions"`
	CreatedAt     mysql.NullTime `json:"created" gorm:"column:created"`
	UpdatedAt     mysql.NullTime `json:"updated" gorm:"column:updated"`
	DeletedAt     mysql.NullTime `json:"deleted" gorm:"column:deleted"` //Soft delete feature
}

//TableName : table name for Gorm
func (q *Questionary) TableName() string {
	return "questionaries"
}

// QuestionaryClient : defines the interface to access Questionary data
type QuestionaryClient interface {
	GetQuestionary(ctx context.Context, QuestionaryID int64) (questionary Questionary, err error)
	GetQuestionaryByLessonID(ctx context.Context, LessonID int64) (questionary Questionary, err error)
	CreateQuestionary(ctx context.Context, questionary *Questionary) error
	UpdateQuestionary(ctx context.Context, questionary *Questionary) error
	DeleteQuestionary(ctx context.Context, questionaryID int64) error
}

// asserts Client implements the QuestionaryClient interface
var _ QuestionaryClient = (*Client)(nil)

func (c Client) GetQuestionary(ctx context.Context, QuestionaryID int64) (questionary Questionary, err error) {
	q := Questionary{}
	err = c.db.Table("Questionary").Where("QuestionaryID = ?", QuestionaryID).Find(&q).Error
	if err != nil {
		fmt.Println(err)
		return q, nil
	}
	return q, nil
}

func (c Client) GetQuestionaryByLessonID(ctx context.Context, LessonID int64) (questionary Questionary, err error) {
	//TODO
	return
}
func (c Client) CreateQuestionary(ctx context.Context, questionary *Questionary) error {
	err := c.db.Create(&questionary).Error
	if err != nil {
		return err
	}
	return nil
}
func (c Client) UpdateQuestionary(ctx context.Context, questionary *Questionary) error {
	//TODO
	return nil
}
func (c Client) DeleteQuestionary(ctx context.Context, questionaryID int64) error {
	//TODO
	return nil
}
