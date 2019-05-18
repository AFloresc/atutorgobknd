package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func getEnv(name string) (value string, err error) {
	value = strings.TrimSpace(os.Getenv(name))

	if value == "" {
		return "", fmt.Errorf("Environment variable %s must be set", name)
	}
	return
}

func getEnvOrDefault(name string, defaultValue string) (value string) {
	value = strings.TrimSpace(os.Getenv(name))

	if value == "" {
		fmt.Printf("Environment variable %s not set, using default value: %s", name, defaultValue)
		return defaultValue
	}

	return
}

func GetGoogleProjectID() (value string, err error) {
	return getEnv("GOOGLE_PROJECT_ID")
}

func GetGoogleStorageBucket() (value string, err error) {
	return getEnv("GOOGLE_STORAGE_BUCKET")
}

func GetMySQLConfig() (dbConfig *mysql.Config, err error) {
	dbConfig = mysql.NewConfig()

	dbConfig.User = getEnvOrDefault("MYSQL_USER", "root")
	if err != nil {
		return
	}

	dbConfig.Passwd = getEnvOrDefault("MYSQL_PASSWORD", "Bautista21")
	if err != nil {
		return
	}

	dbConfig.Addr = getEnvOrDefault("MYSQL_HOSTNAME", "localhost:3306")
	if err != nil {
		return
	}

	dbConfig.DBName = getEnvOrDefault("MYSQL_DBNAME", "atutor_dev")
	if err != nil {
		return
	}

	dbConfig.Collation = getEnvOrDefault("MYSQL_COLLATION", "utf8mb4_unicode_ci")
	dbConfig.ParseTime = true
	dbConfig.Net = "tcp"
	return
}
