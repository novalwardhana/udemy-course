// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package simple

// Injectors from injector.go:

func InitializedUserService(log bool) *UserService {
	userRepository := NewUserRepository(log)
	databasePostgreSQL := NewDatabasePostgreSQL()
	databaseMySQL := NewDatabaseMySQL()
	databaseOracle := NewDatabaseOracle()
	userService := NewUserService(userRepository, databasePostgreSQL, databaseMySQL, databaseOracle)
	return userService
}
