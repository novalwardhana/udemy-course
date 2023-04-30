package simple

/* Example 1 */
type PostgreSQL struct {
	Driver string
	Name   string
}

type MySQL struct {
	Driver string
	Name   string
}

type MongoDB struct {
	Driver string
	Name   string
}

type Database struct {
	pg    *PostgreSQL
	mysql *MySQL
	mongo *MongoDB
}

func NewDatabase(pg *PostgreSQL, mysql *MySQL, mongo *MongoDB) *Database {
	return &Database{
		pg:    pg,
		mysql: mysql,
		mongo: mongo,
	}
}

/* Example 2 */
type Category interface {
	GetData()
	Create()
	Update()
	Delete()
}

type CategoryImpl struct {
}

func (c *CategoryImpl) GetData() {
}

func (c *CategoryImpl) Create() {
}

func (c *CategoryImpl) Update() {
}

func (c *CategoryImpl) Delete() {
}

/* Example 3 */
type UserRepository interface {
	// GetData()
	// Create()
	// Update()
	// Delete()
}
type UserRepositoryImpl struct {
}

func (ur *UserRepositoryImpl) GetData() {
}
func (ur *UserRepositoryImpl) Create() {
}
func (ur *UserRepositoryImpl) Update() {
}
func (ur *UserRepositoryImpl) Delete() {
}

type UserService interface {
}
type UserServiceImpl struct {
	UserRepository UserRepository
}

func (us *UserServiceImpl) GetData() {
}
func (us *UserServiceImpl) Create() {
}
func (us *UserServiceImpl) Update() {
}
func (us *UserServiceImpl) Delete() {
}

type UserController interface {
	GetData()
	Create()
	Update()
	Delete()
}
type UserControllerImpl struct {
	UserService UserService
}

func (uc *UserControllerImpl) GetData() {
}
func (uc *UserControllerImpl) Create() {
}
func (uc *UserControllerImpl) Update() {
}
func (uc *UserControllerImpl) Delete() {
}
