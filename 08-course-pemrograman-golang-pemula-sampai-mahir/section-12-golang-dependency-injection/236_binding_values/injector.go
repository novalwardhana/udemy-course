//go:build wireinject

// + wireinject

package simple

import (
	"io"
	"os"

	"github.com/google/wire"
)

var pg = &PostgreSQL{Driver: "postgresql"}
var mysql = &MySQL{Driver: "mysql"}
var mongo = &MongoDB{Driver: "mongo"}

func InitializationDatabase1() *Database {
	wire.Build(
		wire.Value(pg),
		wire.Value(mysql),
		wire.Value(mongo),
		NewDatabase,
	)
	return nil
}

func InitializationDatabase2() *Database {
	wire.Build(
		wire.Value(&PostgreSQL{Driver: "postgresql"}),
		wire.Value(&MySQL{Driver: "mysql"}),
		wire.Value(&MongoDB{Driver: "mongo"}),
		NewDatabase,
	)
	return nil
}

func InitializedBindingInterface() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

func InitializedCategory() Category {
	wire.Build(wire.InterfaceValue(new(Category), &CategoryImpl{}))
	return nil
}

var userService = wire.NewSet(
	wire.InterfaceValue(
		new(UserService),
		&UserServiceImpl{},
	),
)

func InitializedUserController() UserController {
	wire.Build(
		wire.InterfaceValue(
			new(UserController), &UserControllerImpl{
				UserService: wire.InterfaceValue(new(UserService), &UserServiceImpl{
					UserRepository: wire.InterfaceValue(new(UserRepository), &UserServiceImpl{}),
				}),
			},
		),
	)
	return nil
}
