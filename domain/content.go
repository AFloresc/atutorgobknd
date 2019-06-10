package domain

import (
	"context"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

//Content :
type Content struct {
	ContentID int64          `json:"contentID" gorm:"column:id;primary_key"`
	Title     string         `json:"title" gorm:"column:title"`
	LessonID  int64          `json:"lessonID" gorm:"column:lesson_id"`
	Content   string         `json:"content" gorm:"column:content;type:text"`
	Image     string         `json:"image" gorm:"column:Image;type:text"`
	Position  int64          `json:"position" gorm:"column:position"`
	CreatedAt mysql.NullTime `json:"created" gorm:"column:created"`
	UpdatedAt mysql.NullTime `json:"updated" gorm:"column:updated"`
	DeletedAt mysql.NullTime `json:"deleted" gorm:"column:deleted"` //Soft delete feature
}

// TableName sets the insert table name for this struct type
func (c *Content) TableName() string {
	return "content"
}

// ContentClient : defines the interface to access Content data
type ContentClient interface {
	GetContent(ctx context.Context, contentID int64) (content Content, err error)
	GetAllContentByLessonID(ctx context.Context, lessonID int64) (contents []Content, err error)
	CreateContent(ctx context.Context, content *Content) error
	UpdateContent(ctx context.Context, content *Content) error
	DeleteContent(ctx context.Context, contentID int64) error
}

// asserts Client implements the ContentClient interface
var _ ContentClient = (*Client)(nil)

// GetContent :
func (c Client) GetContent(ctx context.Context, contentID int64) (content Content, err error) {
	ct := Content{}
	err = c.db.Table("content").Where("ContentID = ?", contentID).Find(&ct).Error
	if err != nil {
		fmt.Println(err)
	}
	return ct, nil
}

// GetAllContentByLessonID :
func (c Client) GetAllContentByLessonID(ctx context.Context, lessonID int64) (contents []Content, err error) {
	cts := []Content{}
	err = c.db.Table("content").Where("lesson_id = ?", lessonID).Find(&cts).Error
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	return cts, nil
}

// CreateContent :
func (c Client) CreateContent(ctx context.Context, content *Content) error {
	err := c.db.Create(&content).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateContent :
func (c Client) UpdateContent(ctx context.Context, content *Content) error {
	err := c.db.Save(content).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteContent :
func (c Client) DeleteContent(ctx context.Context, contentID int64) error {
	if contentID == 0 {
		return fmt.Errorf("Error!!! (DeleteContent), incorrect Content ID: %d", contentID)
	}

	content := Content{
		ContentID: contentID,
	}
	return c.db.Delete(&content).Error
}
