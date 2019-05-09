package domain

import (
	"context"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// Concept defines the repository of several concepts related to Android
type Concept struct {
	ConceptID   int64          `json:"conceptID" gorm:"column:id;primary_key" `
	Title       string         `json:"title" gorm:"column:title" `
	Description string         `json:"description" gorm:"column:description;type:text" `
	Language    string         `json:"language" gorm:"column:language" `
	CreatedAt   mysql.NullTime `json:"created" gorm:"column:created"`
	UpdatedAt   mysql.NullTime `json:"updated" gorm:"column:updated"`
	DeletedAt   mysql.NullTime `json:"deleted" gorm:"column:deleted"` //Soft delete feature
}

// TableName sets the insert table name for this struct type
func (c *Concept) TableName() string {
	return "concept"
}

// ConceptClient defines the interface to access ProfilesUser data
type ConceptClient interface {
	GetConcept(ctx context.Context, conceptID int64) (concept Concept, err error)
	CreateConcept(ctx context.Context, concept *Concept) error
	UpdateConcept(ctx context.Context, concept *Concept) error
	GetConceptBySearch(ctx context.Context, search string) (concepts []Concept, err error)
	DeleteConcept(ctx context.Context, conceptID int64) error
}

// asserts Client implements the ConceptClient interface
var _ ConceptClient = (*Client)(nil)

func (c Client) GetConcept(ctx context.Context, conceptID int64) (concept Concept, err error) {
	//TODO
	return concept, nil
}
func (c Client) CreateConcept(ctx context.Context, concept *Concept) error {
	//TODO
	return nil
}
func (c Client) GetConceptBySearch(ctx context.Context, search string) (concepts []Concept, err error) {
	//TODO
	concepts = []Concept{}

	err = c.db.Table("Concept").Select("conceptID, title, description,").Where("title LIKE %?%", search).Scan(&concepts).Error
	if err != nil {
		return nil, err
	}
	return concepts, nil

}

// UpdateConcept : updates a Concept
func (c Client) UpdateConcept(ctx context.Context, concept *Concept) error {
	err := c.db.Save(concept).Error
	if err != nil {
		return err
	}
	return nil
}

func (c Client) DeleteConcept(ctx context.Context, conceptID int64) error {
	if conceptID == 0 {
		return fmt.Errorf("Error!!! (DeleteConcept), incorrect Concept ID: %d", conceptID)
	}

	concept := Concept{
		ConceptID: conceptID,
	}
	return c.db.Delete(&concept).Error

}
