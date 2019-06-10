package domain

import (
	"context"
	"fmt"

	"github.com/atutor/utils"
	"github.com/go-sql-driver/mysql"
)

//Lesson represent every box inside the app
type Lesson struct {
	LessonID    int64          `json:"lessonID" gorm:"column:lessonID;primary_key"`
	Title       string         `json:"titlee" gorm:"column:Title"`
	Text        string         `json:"text" gorm:"column:Text;type:text"`
	Image       string         `json:"image" gorm:"column:Image;type:text"`
	Language    string         `json:"language" gorm:"column:Language"`
	Contents    []Content      `json:"contents"`
	Questionary Questionary    `json:"questionary"`
	CourseID    int64          `json:"courseID" gorm:"column:courseID"`
	CreatedAt   mysql.NullTime `json:"created" gorm:"column:created"`
	UpdatedAt   mysql.NullTime `json:"updated" gorm:"column:updated"`
	DeletedAt   mysql.NullTime `json:"deleted" gorm:"column:deleted"` //Soft delete feature
	Position    int64          `json:"position" gorm:"column:position"`
}

//TableName : table name for Gorm
func (l *Lesson) TableName() string {
	return "lesson"
}

// LessonClient : defines the interface to access Lesson data
type LessonClient interface {
	GetLesson(ctx context.Context, lessonID int64) (lesson *Lesson, err error)
	GetAllLessonsByCourseID(ctx context.Context, courseID int64) (lessons []Lesson, err error)
	CreateLesson(ctx context.Context, lesson *Lesson) error
	UpdateLesson(ctx context.Context, lesson *Lesson) error
	DeleteLesson(ctx context.Context, lessonID int64) error
}

// asserts Client implements the LessonClient interface
var _ LessonClient = (*Client)(nil)

// GetLesson : retrieves a lesson by it's ID
func (c Client) GetLesson(ctx context.Context, lessonID int64) (lesson *Lesson, err error) {
	ls := Lesson{}
	err = c.db.Table("lesson").Where("lessonID = ?", lessonID).Find(&ls).Error
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	contents, err := c.GetAllContentByLessonID(ctx, ls.LessonID)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	for _, content := range contents {
		ls.Contents = append(ls.Contents, content)
	}
	return &ls, nil
}
func (c Client) CreateLesson(ctx context.Context, lesson *Lesson) error {
	err := c.db.Create(&lesson).Error
	if err != nil {
		return err
	}
	return nil
}
func (c Client) UpdateLesson(ctx context.Context, lesson *Lesson) error {
	err := c.db.Save(&lesson).Error
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

func (c Client) GetAllLessonsByCourseID(ctx context.Context, courseID int64) (lessons []Lesson, err error) {

	ls := []Lesson{}
	err = c.db.Table("lessons").Where("courseID = ?", courseID).Find(&ls).Error
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	for index, lesson := range ls {
		contents, err := c.GetAllContentByLessonID(ctx, lesson.LessonID)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, content := range contents {
			ls[index].Contents = append(ls[index].Contents, content)
		}
	}
	return ls, nil
}

// GetLessonsByLanguage : retrieves all the concepts in a language
func (c Client) GetLessonsByLanguage(ctx context.Context, language string) (lessons []Lesson, err error) {
	lessns := []Lesson{}
	if !utils.ValidateLanguage(language) {
		return nil, utils.NewError("Unknown Language.")
	}
	err = c.db.Table("lesson").Where("language = ?", language).Find(&lessns).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	for index, lesson := range lessns {
		contents, err := c.GetAllContentByLessonID(ctx, lesson.LessonID)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, content := range contents {
			lessns[index].Contents = append(lessns[index].Contents, content)
		}
	}
	return lessns, nil
}

func (c Client) hardDeleteLesson(ctx context.Context, lessonID int64) error {
	if lessonID != 0 {
		return c.db.Exec("DELETE FROM lessons WHERE lessonID=? ", lessonID).Error
	}
	return utils.NewError("Lesson ID value not allowed on hard Delete action")
}
