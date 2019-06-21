package domain

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
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
		&Question{},
		&UserLessons{},
		&Mark{},
		&Questionary{}).Error
}

// Transaction :
func (s *Client) Transaction() (*Client, error) {
	client := Client{
		db: s.db.Begin(),
	}
	return &client, nil
}

// NewTransaction : This method starts a new transaction
func NewTransaction(s *Client) (*Client, error) {
	client := Client{
		db: s.db.Begin(),
	}
	return &client, nil
}

// Rollback : This method makes a rollback on transaction
func (s *Client) Rollback() {
	s.db.Rollback()
}

// Commit : This method commits transaction
func (s *Client) Commit() error {
	return s.db.Commit().Error
}

// Close close current db connection.  If database connection is not an io.Closer, returns an error.
func (s *Client) Close() error {
	return s.db.Close()
}

// Initialize : Initializes the mysql database connection
func (s *Client) Initialize(user string, password string, ip string, port int, dbname string) {
	//connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, ip, port, dbname)
	// println("Connection: ", connectionString)
	myInstanceName := "atutor:europe-west1:atutor-develop-db"
	connectionString := fmt.Sprintf("%s:%s@cloudsql(%s)/%s?charset=utf8&parseTime=True&loc=UTC", user, password, myInstanceName, dbname)
	//dbURI := fmt.Sprintf("%s:%s@unix(%s/%s)/", user, password, "/cloudsql", myInstanceName)

	var err error
	s.db, err = gorm.Open("mysql", connectionString)
	//s.db, err = gorm.Open("mysql", dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection Established.")
	}
}
