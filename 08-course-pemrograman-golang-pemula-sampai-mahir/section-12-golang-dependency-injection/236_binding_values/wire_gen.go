// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

// Injectors from injector.go:

func InitializationDatabase1() *Database {
	postgreSQL := _wirePostgreSQLValue
	mySQL := _wireMySQLValue
	mongoDB := _wireMongoDBValue
	database := NewDatabase(postgreSQL, mySQL, mongoDB)
	return database
}

var (
	_wirePostgreSQLValue = pg
	_wireMySQLValue      = mysql
	_wireMongoDBValue    = mongo
)

func InitializationDatabase2() *Database {
	postgreSQL := _wireSimplePostgreSQLValue
	mySQL := _wireSimpleMySQLValue
	mongoDB := _wireSimpleMongoDBValue
	database := NewDatabase(postgreSQL, mySQL, mongoDB)
	return database
}

var (
	_wireSimplePostgreSQLValue = &PostgreSQL{Driver: "postgresql"}
	_wireSimpleMySQLValue      = &MySQL{Driver: "mysql"}
	_wireSimpleMongoDBValue    = &MongoDB{Driver: "mongo"}
)

func InitializedBindingInterface() io.Reader {
	reader := _wireFileValue
	return reader
}

var (
	_wireFileValue = os.Stdin
)

func InitializedCategory() Category {
	category := _wireCategoryImplValue
	return category
}

var (
	_wireCategoryImplValue = &CategoryImpl{}
)

func InitializedUserController() UserController {
	userController := _wireUserControllerImplValue
	return userController
}

var (
	_wireUserControllerImplValue = &UserControllerImpl{
		UserService: wire.InterfaceValue(new(UserService), &UserServiceImpl{
			UserRepository: wire.InterfaceValue(new(UserRepository), &UserServiceImpl{}),
		}),
	}
)

// injector.go:

var pg = &PostgreSQL{Driver: "postgresql"}

var mysql = &MySQL{Driver: "mysql"}

var mongo = &MongoDB{Driver: "mongo"}

var userService = wire.NewSet(wire.InterfaceValue(
	new(UserService),
	&UserServiceImpl{},
),
)