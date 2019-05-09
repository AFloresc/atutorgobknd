package domain

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

// UserProfile defines the table for managing user access to profiles
type User struct {
	UserID    int64          `gorm:"column:id;primary_key" json:"userID"`
	Name      string         `gorm:"column:name" json:"name"`
	Password  string         `gorm:"column:password" json:"password"`
	Email     string         `gorm:"column:email" json:"email"`
	Role      string         `gorm:"-" json:"role"`
	CreatedAt mysql.NullTime `json:"created" gorm:"column:created"`
	UpdatedAt mysql.NullTime `json:"updated" gorm:"column:updated"`
	DeletedAt mysql.NullTime `json:"deleted" gorm:"column:deleted"` //Soft delete feature
}

// TableName sets the insert table name for this struct type
func (p *User) TableName() string {
	return "user"
}

// UserProfilesClient defines the interface to access ProfilesUser data
type UserClient interface {
	GetUserByPassword(ctx context.Context, userID int64) (user User, err error)
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (user User, err error)
	GetUserProfiles(ctx context.Context, userID int64) ([]User, error)
}

// asserts Client implements the UserProfilesClient interface
var _ UserClient = (*Client)(nil)

// Creates a User into database
func (s Client) CreateUser(ctx context.Context, user *User) error {
	return s.db.Create(user).Error
}

func (s Client) GetUserByPassword(ctx context.Context, userID int64) (user User, err error) {

	err = s.db.Where("user_id = ?", userID).Find(&user).Error
	if err != nil {
		return
	}

	return
}

func (s Client) GetUserByEmail(ctx context.Context, email string) (user User, err error) {

	err = s.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return
	}

	return
}

func (s Client) GetUserProfiles(ctx context.Context, userID int64) (profiles []User, err error) {
	profiles = []User{}
	err = s.db.Where("user_id = ?", userID).Find(&profiles).Error
	if err != nil {
		return
	}
	return
}
