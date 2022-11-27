package config

import "fmt"

const (
	DBHost     = "serverdb"
	DBPort     = "3306"
	DBUser     = "mydbuser"
	DBPassword = "mydbpassword"
	DbName     = "mydatabase"
)

func GetDBConnection() string {
	connection := fmt.Sprintf("%s:%s:@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=Tre&loc=Local",
		DBUser,
		DBPassword,
		DBHost,
		DBPort,
		DbName)

	return connection
}
