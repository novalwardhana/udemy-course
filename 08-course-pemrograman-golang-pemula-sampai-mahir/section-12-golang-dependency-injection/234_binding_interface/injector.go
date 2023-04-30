//go:build wireinject

// + wireinject

package simple

import "github.com/google/wire"

var bindUserRepository = wire.NewSet(
	NewUserRepository,
	wire.Bind(new(UserRepository), new(*UserRepositoryImpl)),
)

func InitializedUserService() *UserServiceImpl {
	wire.Build(bindUserRepository, NewUserService)
	return nil
}

var bindCategoryRepository = wire.NewSet(
	NewCategoryRepository,
	wire.Bind(new(CategoryRepository), new(*CategoryRepositoryImpl)),
)

var bindClassRepository = wire.NewSet(
	NewClassRepository,
	wire.Bind(new(ClassRepository), new(*ClassRepositoryImpl)),
)

func InitializedCategoryClass(log bool) *CategoryClassServiceImpl {
	wire.Build(bindCategoryRepository, bindClassRepository, NewCategoryClassService)
	return nil
}

var bindAdminRepository = wire.NewSet(
	NewAdminRepository,
	wire.Bind(new(AdminRepository), new(*AdminRepositoryImpl)),
)
var bindAdminService = wire.NewSet(
	NewAdminService,
	wire.Bind(new(AdminService), new(*AdminServiceImpl)),
)

func InitializedAdminController(log bool) AdminController {
	wire.Build(bindAdminRepository, bindAdminService, NewAdminController)
	return nil
}
