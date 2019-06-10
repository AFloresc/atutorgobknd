package domain

import (
	"context"
	"fmt"

	"github.com/atutor/utils"
	"github.com/go-sql-driver/mysql"
)

// Course represent every course inside the app
type Course struct {
	CourseID    int64          `json:"courseID" gorm:"column:courseID;primary_key"`
	Name        string         `json:"titlee" gorm:"column:title"`
	Description string         `json:"description" gorm:"column:description;type:text"`
	Target      string         `json:"target" gorm:"column:target;type:text"`
	CreatedAt   mysql.NullTime `json:"created" gorm:"column:created"`
	UpdatedAt   mysql.NullTime `json:"updated" gorm:"column:updated"`
	DeletedAt   mysql.NullTime `json:"deleted" gorm:"column:deleted"` //Soft delete feature
	Lessons     []Lesson       `json:"lessons"`
}

// TableName sets the insert table name for this struct type
func (c *Course) TableName() string {
	return "course"
}

// CourseClient : defines the interface to access Course data
type CourseClient interface {
	GetCourse(ctx context.Context, courseID int64) (course Course, err error)
	GetWholeCourse(ctx context.Context, courseID int64) (course Course, err error)
	CreateCourse(ctx context.Context, course *Course) error
	UpdateCourse(ctx context.Context, course *Course) error
	DeleteCourse(ctx context.Context, courseID int64) error
}

// asserts Client implements the TestClient interface
var _ CourseClient = (*Client)(nil)

//GetCourse : Gets a Course by course ID
func (c Client) GetCourse(ctx context.Context, courseID int64) (course Course, err error) {
	crs := Course{}
	err = c.db.Table("course").Where("CourseID = ?", courseID).Find(&crs).Error
	if err != nil {
		fmt.Println(err)
		return course, nil
	}
	lessons, err := c.GetAllLessonsByCourseID(ctx, crs.CourseID)
	if err != nil {
		fmt.Println(err)
		return course, nil
	}

	for _, lesson := range lessons {
		crs.Lessons = append(crs.Lessons, lesson)

	}
	return crs, nil
}

//GetWholeCourse : Gets a Course by course ID
func (c Client) GetWholeCourse(ctx context.Context, courseID int64) (course Course, err error) {
	crs := Course{}
	err = c.db.Table("course").Where("CourseID = ?", courseID).Find(&crs).Error
	if err != nil {
		fmt.Println(err)
		return course, nil
	}
	return crs, nil
}

//CreateCourse : Creates a Coure by Course object
func (c Client) CreateCourse(ctx context.Context, course *Course) error {
	err := c.db.Create(&course).Error
	if err != nil {
		return err
	}

	return nil
}

// UpdateCourse :
func (c Client) UpdateCourse(ctx context.Context, course *Course) error {
	err := c.db.Save(&course).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteCourse :
func (c Client) DeleteCourse(ctx context.Context, courseID int64) error {
	if courseID == 0 {
		return fmt.Errorf("Error!!! (DeleteCurse), incorrect Concept ID: %d", courseID)
	}
	course := Course{
		CourseID: courseID,
	}
	return c.db.Delete(&course).Error
}

// hardDeleteCurse : deletes permanently a course, WARNING!!! for testing purposes only!!!
func (c Client) hardDeleteCourse(ctx context.Context, courseID int64) error {
	if courseID != 0 {
		return c.db.Exec("DELETE FROM course WHERE courseID=? ", courseID).Error
	}
	return utils.NewError("Course ID value not allowed on hard Delete action")
}
