package domain

import "context"

//Lesson represent every box inside the app
type Course struct {
	CourseID    int64  `json:"courseID" gorm:"column:courseID;primary_key"`
	Name        string `json:"titlee" gorm:"column:Title"`
	Description string `json:"description" gorm:"column:Description;type:text"`
}

// TableName sets the insert table name for this struct type
func (c *Course) TableName() string {
	return "course"
}

// CourseClient : defines the interface to access Course data
type CourseClient interface {
	GetCourse(ctx context.Context, courseID int64) (Test *Test, err error)
	CreateCourse(ctx context.Context, course *Course) error
	UpdateCourse(ctx context.Context, course *Course) error
	DeleteCourse(ctx context.Context, courseID int64) error
}

// asserts Client implements the TestClient interface
var _ CourseClient = (*Client)(nil)

//GetCourse : Gets a Course by course ID
func (c Client) GetCourse(ctx context.Context, courseID int64) (Test *Test, err error) {
	//TODO

	return nil, nil
}

//CreateCourse : Creates a Coure by Course object
func (c Client) CreateCourse(ctx context.Context, course *Course) error {
	//TODO

	return nil
}

func (c Client) UpdateCourse(ctx context.Context, course *Course) error {
	//TODO

	return nil
}

func (c Client) DeleteCourse(ctx context.Context, courseID int64) error {
	//TODO

	return nil
}
