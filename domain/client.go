package domain

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Client handles the db connection, performs basic operation to websays mysql tables
type Client struct {
	db *gorm.DB
}

// NewClient initialize a new db connection, needs a valid mysql.Config
func NewClient(config *mysql.Config) (*Client, error) {
	config.ParseTime = true
	db, err := gorm.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	client := Client{
		db: db,
	}
	return &client, nil
}

// AutoMigrate run auto migration for given models, will only add missing fields, won't delete/change current data
func (s *Client) AutoMigrate() error {

	return s.db.AutoMigrate(
		&Course{},
		&Lesson{},
		&Content{},
		&User{},
		&Concept{},
		&Content{},
		&Question{}).Error
}

func (s *Client) Transaction() (*Client, error) {
	client := Client{
		db: s.db.Begin(),
	}
	return &client, nil
}

func NewTransaction(s *Client) (*Client, error) {
	client := Client{
		db: s.db.Begin(),
	}
	return &client, nil
}

func (s *Client) Rollback() {
	s.db.Rollback()
}

func (s *Client) Commit() error {
	return s.db.Commit().Error
}

// Close close current db connection.  If database connection is not an io.Closer, returns an error.
func (s *Client) Close() error {
	return s.db.Close()
}

func (s *Client) Initialize(user string, password string, ip string, port int, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, ip, port, dbname)
	println("Connection: ", connectionString)
	var err error
	s.db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection Established.")
	}
}
