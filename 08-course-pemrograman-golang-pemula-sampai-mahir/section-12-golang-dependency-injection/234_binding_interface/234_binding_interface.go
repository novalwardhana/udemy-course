package simple

import "fmt"

/* Example 1 */

type UserRepository interface {
	Message(name string) string
}

type UserRepositoryImpl struct {
}

func (u *UserRepositoryImpl) Message(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

type UserServiceImpl struct {
	UserRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

/* Example 2 */
type CategoryRepository interface {
	GetData()
	Create()
	Update()
	Delete()
}

type CategoryRepositoryImpl struct {
	Log bool
}

func (c *CategoryRepositoryImpl) GetData() {
}

func (c *CategoryRepositoryImpl) Create() {
}

func (c *CategoryRepositoryImpl) Update() {
}

func (c *CategoryRepositoryImpl) Delete() {
}

func NewCategoryRepository(log bool) *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{
		Log: log,
	}
}

type ClassRepository interface {
	GetData()
	Create()
	Update()
	Delete()
}

type ClassRepositoryImpl struct {
	Log bool
}

func (c *ClassRepositoryImpl) GetData() {
}

func (c *ClassRepositoryImpl) Create() {
}

func (c *ClassRepositoryImpl) Update() {
}

func (c *ClassRepositoryImpl) Delete() {
}

func NewClassRepository(log bool) *ClassRepositoryImpl {
	return &ClassRepositoryImpl{
		Log: log,
	}
}

type CategoryClassServiceImpl struct {
	CategoryRepository CategoryRepository
	ClassRepository    ClassRepository
}

func NewCategoryClassService(categoryRepository CategoryRepository, classRepository ClassRepository) *CategoryClassServiceImpl {
	return &CategoryClassServiceImpl{
		CategoryRepository: categoryRepository,
		ClassRepository:    classRepository,
	}
}

/* Example 3 */
type AdminRepository interface {
	GetData()
	UploadData()
}

type AdminRepositoryImpl struct {
	Log bool
}

func (a *AdminRepositoryImpl) GetData() {
}

func (a *AdminRepositoryImpl) UploadData() {
}

func NewAdminRepository(log bool) *AdminRepositoryImpl {
	return &AdminRepositoryImpl{
		Log: log,
	}
}

type AdminService interface {
	GetData()
	UploadData()
}

type AdminServiceImpl struct {
	AdminRepository AdminRepository
}

func (a *AdminServiceImpl) GetData() {
	a.AdminRepository.GetData()
}

func (a *AdminServiceImpl) UploadData() {
	a.AdminRepository.UploadData()
}

func NewAdminService(adminRepository AdminRepository) *AdminServiceImpl {
	return &AdminServiceImpl{
		AdminRepository: adminRepository,
	}
}

type AdminController interface {
	GetData()
	UploadData()
}

type AdminControllerImpl struct {
	AdminService AdminService
}

func (ac *AdminControllerImpl) GetData() {
	ac.AdminService.GetData()
}

func (ac *AdminControllerImpl) UploadData() {
	ac.AdminService.UploadData()
}

func NewAdminController(adminService AdminService) AdminController {
	return &AdminControllerImpl{
		AdminService: adminService,
	}
}
