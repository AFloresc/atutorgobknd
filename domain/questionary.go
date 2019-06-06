package domain

import (
	"context"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

//Questionary represent every box inside the app
type Questionary struct {
	QuestionaryID int64          `json:"questionaryID" gorm:"column:questionaryID;primary_key"`
	Description   string         `json:"description" gorm:"column:description"`
	LessonID      int64          `json:"lessonID" gorm:"column:lessonID"`
	Questions     []Question     `json:"questions"`
	CreatedAt     mysql.NullTime `json:"created" gorm:"column:created"`
	UpdatedAt     mysql.NullTime `json:"updated" gorm:"column:updated"`
	DeletedAt     mysql.NullTime `json:"deleted" gorm:"column:deleted"` //Soft delete feature
}

//TableName : table name for Gorm
func (q *Questionary) TableName() string {
	return "questionary"
}

// QuestionaryClient : defines the interface to access Questionary data
type QuestionaryClient interface {
	GetQuestionary(ctx context.Context, questionaryID int64) (questionary Questionary, err error)
	GetQuestionaryByLessonID(ctx context.Context, LessonID int64) (questionary Questionary, err error)
	CreateQuestionary(ctx context.Context, questionary *Questionary) error
	UpdateQuestionary(ctx context.Context, questionary *Questionary) error
	DeleteQuestionary(ctx context.Context, questionaryID int64) error
}

// asserts Client implements the QuestionaryClient interface
var _ QuestionaryClient = (*Client)(nil)

// GetQuestionary :
func (c Client) GetQuestionary(ctx context.Context, questionaryID int64) (questionary Questionary, err error) {
	q := Questionary{}
	err = c.db.Table("Questionary").Where("QuestionaryID = ?", questionaryID).Find(&q).Error
	if err != nil {
		fmt.Println(err)
		return q, nil
	}
	questions, err := c.GetQuestionsByQuestionaryID(ctx, questionaryID)
	if err != nil {
		fmt.Println(err)
		return q, nil
	}

	for _, question := range questions {
		q.Questions = append(q.Questions, question)
	}

	return q, nil
}

// GetQuestionaryByLessonID :
func (c Client) GetQuestionaryByLessonID(ctx context.Context, LessonID int64) (questionary Questionary, err error) {
	q := Questionary{}
	err = c.db.Table("Questionary").Where("LessonID = ?", LessonID).Find(&q).Error
	if err != nil {
		fmt.Println(err)
		return q, nil
	}
	questions, err := c.GetQuestionsByQuestionaryID(ctx, q.QuestionaryID)
	if err != nil {
		fmt.Println(err)
		return q, nil
	}

	for _, question := range questions {
		q.Questions = append(q.Questions, question)
	}

	return q, nil
}

// CreateQuestionary :
func (c Client) CreateQuestionary(ctx context.Context, questionary *Questionary) error {
	err := c.db.Create(&questionary).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateQuestionary :
func (c Client) UpdateQuestionary(ctx context.Context, questionary *Questionary) error {
	err := c.db.Save(&questionary).Error
	if err != nil {
		return err
	}
	return nil
}

//  DeleteQuestionary :
func (c Client) DeleteQuestionary(ctx context.Context, questionaryID int64) error {
	if questionaryID == 0 {
		return fmt.Errorf("Error!!! (DeleteQuestionary), incorrect Questionary ID: %d", questionaryID)
	}
	questionary := Questionary{
		QuestionaryID: questionaryID,
	}
	return c.db.Delete(&questionary).Error
}
