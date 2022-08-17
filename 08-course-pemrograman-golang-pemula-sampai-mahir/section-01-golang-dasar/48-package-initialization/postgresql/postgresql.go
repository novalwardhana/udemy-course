package postgresql

import "fmt"

var database string
var port string
var host string

func init() {
	fmt.Println("init postgresql")
	database = "Postgresql"
	port = "5432"
	host = "localhost"
}

func GetDatabase() string {
	return database
}

func GetPort() string {
	return port
}

func GetHost() string {
	return host
}
