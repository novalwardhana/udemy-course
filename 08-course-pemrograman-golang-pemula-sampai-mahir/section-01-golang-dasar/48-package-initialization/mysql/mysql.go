package mysql

import "fmt"

var database string
var port string
var host string

func init() {
	fmt.Println("init mysql")
	database = "mysql"
	port = "3306"
	host = "localhost"
}
