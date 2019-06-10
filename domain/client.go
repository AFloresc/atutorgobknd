package domain

import (
	"fmt"
	"log"
	"os"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
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
	env := os.Getenv("ATUTOR_ENV")
	connectionString := ""
	if env == "test" {
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, ip, port, dbname)
	} else {
		myInstanceName := "atutor:europe-west1:atutor-develop-db"
		connectionString = fmt.Sprintf("%s:%s@cloudsql(%s)/%s?charset=utf8&parseTime=True&loc=UTC", user, password, myInstanceName, dbname)
	}
	//connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, password, ip, port, dbname)
	//sql.Open("mysql", "cloudsql:my-instance*dbname/user/passwd")

	// connectionStringDev := fmt.Sprintf("cloudsql:%s/%s/%s", myInstanceName, user, password)
	// var dbConnectionString = fmt.Sprintf("%s:%s@cloudsql(atutor-develop-db)/%?charset=utf8&parseTime=True&loc=UTC", user, password
	// var dbConnectionString = fmt.Sprintf("%s:%s@cloudsql(%s)/%s?charset=utf8&parseTime=True&loc=UTC", user, password, myInstanceName, dbname)
	//user:password@cloudsql(copiedPastedInstanceConnectionName)/databaseName?charset=charset&collation=collation&tls=tlsConfigName&parseTime=true
	println("Connection: ", connectionString)
	var err error
	s.db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection Established.")
	}

}
