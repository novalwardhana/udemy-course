//go:build wireinject

// + wireinject

package simple

import "github.com/google/wire"

func InitializedUserController(name string) (*UserController, func()) {
	wire.Build(NewUserRepository, NewUserService, NewUserController)
	return nil, nil
}

var adminRepository = wire.NewSet(
	NewAdminRepository,
	wire.Bind(new(AdminRepository), new(*AdminRepositoryImpl)),
)

var adminService = wire.NewSet(
	NewAdminService,
	wire.Bind(new(AdminService), new(*AdminServiceImpl)),
)

func InitializedAdminController(name string) (*AdminControllerImpl, func(), error) {
	wire.Build(adminRepository, adminService, NewAdminController)
	return nil, nil, nil
}
