package domain

import (
	"context"

	"github.com/go-sql-driver/mysql"
)

//Test represent every box inside the app
type Test struct {
	TestID      int64          `json:"testID" gorm:"column:testID;primary_key"`
	Description string         `json:"description" gorm:"column:description"`
	LessonID    int64          `json:"TestID" gorm:"column:Test_id"`
	Questions   []Question     `json:"questions"`
	CreatedAt   mysql.NullTime `json:"created" gorm:"column:created"`
	UpdatedAt   mysql.NullTime `json:"updated" gorm:"column:updated"`
	DeletedAt   mysql.NullTime `json:"deleted" gorm:"column:deleted"` //Soft delete feature
}

//TableName : table name for Gorm
func (t *Test) TableName() string {
	return "tests"
}

// TestClient : defines the interface to access Test data
type TestClient interface {
	GetTest(ctx context.Context, testID int64) (Test *Test, err error)
	GetTestByTestID(ctx context.Context, TestID int64) (test Test, err error)
	CreateTest(ctx context.Context, test *Test) error
	UpdateTest(ctx context.Context, test *Test) error
	DeleteTest(ctx context.Context, testID int64) error
}

// asserts Client implements the TestClient interface
var _ TestClient = (*Client)(nil)

func (c Client) GetTest(ctx context.Context, testID int64) (Test *Test, err error) {
	//TODO
	return nil, nil
}

func (c Client) GetTestByTestID(ctx context.Context, TestID int64) (test Test, err error) {
	//TODO
	return
}
func (c Client) CreateTest(ctx context.Context, test *Test) error {
	//TODO
	return nil
}
func (c Client) UpdateTest(ctx context.Context, test *Test) error {
	//TODO
	return nil
}
func (c Client) DeleteTest(ctx context.Context, testID int64) error {
	//TODO
	return nil
}
