//go:build wireinject

// + wireinject

package simple

import "github.com/google/wire"

func InitializedUserService(log bool) *UserService {
	wire.Build(NewDatabasePostgreSQL, NewDatabaseMySQL, NewDatabaseOracle, NewUserRepository, NewUserService)
	return nil
}
