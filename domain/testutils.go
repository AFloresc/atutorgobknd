package domain

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func getEnvOrDefault(name string, defaultValue string) (value string) {
	value = os.Getenv(name)

	if value == "" {
		fmt.Printf("Environment variable %s not set, using default value: %s\n", name, defaultValue)
		return defaultValue
	}

	return
}

func NewTestClientFromEnv() (*Client, error) {
	dbConfig := mysql.NewConfig()
	dbConfig.User = getEnvOrDefault("MYSQL_USER", "root")
	dbConfig.Passwd = getEnvOrDefault("MYSQL_PASSWORD", "Bautista21")
	dbConfig.Net = "tcp"
	dbConfig.Addr = getEnvOrDefault("MYSQL_HOSTNAME", "localhost:3306")
	dbConfig.DBName = getEnvOrDefault("MYSQL_DBNAME", "atutor_dev")
	dbConfig.Collation = getEnvOrDefault("MYSQL_COLLATION", "utf8mb4_unicode_ci")

	client, err := NewClient(dbConfig)

	if err != nil {
		return nil, err
	}
	err = client.AutoMigrate()
	if err != nil {
		return nil, err
	}
	return client, nil
}
