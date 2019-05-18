package domain

import (
	"context"
	"fmt"

	"github.com/atutor/utils"
	"github.com/go-sql-driver/mysql"
)

// Concept defines the repository of several concepts related to Android
type Concept struct {
	ConceptID   int64          `json:"conceptID" gorm:"column:id;primary_key" `
	Title       string         `json:"title" gorm:"column:title" `
	Description string         `json:"description" gorm:"column:description;type:text" `
	Language    string         `json:"language" gorm:"column:language" `
	CourseID    int64          `json:"course" gorm:"column:course" `
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
	GetConcept(ctx context.Context, conceptID int64, language string) (concept Concept, err error)
	CreateConcept(ctx context.Context, concept *Concept) error
	UpdateConcept(ctx context.Context, concept *Concept) error
	GetConceptBySearch(ctx context.Context, search string) (concepts []Concept, err error)
	DeleteConcept(ctx context.Context, conceptID int64) error
}

// asserts Client implements the ConceptClient interface
var _ ConceptClient = (*Client)(nil)

func (c Client) GetConcept(ctx context.Context, conceptID int64, language string) (concept Concept, err error) {
	conc := Concept{}
	err = c.db.Table("concept").Where("ID = ? AND language = ?", conceptID, language).Find(&conc).Error
	if err != nil {
		fmt.Println(err)
		return concept, nil
	}
	return conc, nil
}

// CreateConcept : Creates a concept
func (c Client) CreateConcept(ctx context.Context, concept *Concept) error {
	return c.db.Create(concept).Error
}

// GetConceptBySearch : Get concepts defined by a search
func (c Client) GetConceptBySearch(ctx context.Context, search string) (concepts []Concept, err error) {
	//TODO
	concepts = []Concept{}
	search = "%" + search + "%"

	err = c.db.Table("Concept").Select("id, title, description").Where("title LIKE ?", search).Find(&concepts).Error
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

// hardDeleteConcept : deletes permanently a concept,  WARNING!!! for testing purposes only!!!
func (c Client) hardDeleteConcept(ctx context.Context, conceptID int64) error {
	if conceptID != 0 {
		return c.db.Exec("DELETE FROM concept WHERE id=? ", conceptID).Error
	}
	return utils.NewError("Condept ID value not allowed on hard Delete action")
}
