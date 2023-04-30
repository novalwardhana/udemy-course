package simple

import (
	"fmt"
)

/* Example 1 */
type UserRepository struct {
	Name string
}

func (ur *UserRepository) CloseRepo() {
	message := fmt.Sprintf("Close repo %s", ur.Name)
	fmt.Println(message)
}

func NewUserRepository(name string) (*UserRepository, func()) {
	userRepository := &UserRepository{
		Name: name,
	}
	return userRepository, func() {
		userRepository.CloseRepo()
	}
}

type UserService struct {
	*UserRepository
}

func (us *UserService) CloseService() {
	message := fmt.Sprintf("Close service %s", us.UserRepository.Name)
	fmt.Println(message)
}

func NewUserService(userRepository *UserRepository) (*UserService, func()) {
	userService := &UserService{
		UserRepository: userRepository,
	}
	return userService, func() {
		userService.CloseService()
	}
}

type UserController struct {
	*UserService
}

func (uc *UserController) CloseController() {
	message := fmt.Sprintf("Close controller %s", uc.UserService.Name)
	fmt.Println(message)
}

func NewUserController(userService *UserService) (*UserController, func()) {
	userController := &UserController{
		UserService: userService,
	}
	return userController, func() {
		userController.CloseController()
	}
}

/* Example 2 */
type AdminRepository interface {
	Close()
}
type AdminRepositoryImpl struct {
	Name string
}

func (ar *AdminRepositoryImpl) Close() {
	message := fmt.Sprintf("Close repo %s", ar.Name)
	fmt.Println(message)
}

func NewAdminRepository(name string) (*AdminRepositoryImpl, func(), error) {
	adminRepositoryImpl := &AdminRepositoryImpl{
		Name: name,
	}
	return adminRepositoryImpl, func() {
		adminRepositoryImpl.Close()
	}, nil
}

type AdminService interface {
	Close()
}
type AdminServiceImpl struct {
	AdminRepository AdminRepository
}

func (as *AdminServiceImpl) Close() {
	message := "Close service "
	fmt.Println(message)
}
func NewAdminService(adminRepository AdminRepository) (*AdminServiceImpl, func(), error) {
	adminRepositoryImpl := &AdminServiceImpl{
		AdminRepository: adminRepository,
	}
	return adminRepositoryImpl, func() {
		adminRepositoryImpl.Close()
	}, nil
}

type AdminController interface {
	Close()
}
type AdminControllerImpl struct {
	AdminService AdminService
}

func (ac *AdminControllerImpl) Close() {
	message := "Close controller"
	fmt.Println(message)
}
func NewAdminController(adminService AdminService) (*AdminControllerImpl, func(), error) {
	adminControllerImpl := &AdminControllerImpl{
		AdminService: adminService,
	}
	return adminControllerImpl, func() {
		adminControllerImpl.Close()
	}, nil
}
