package simple

type Database struct {
	Driver string
	Uri    string
}

type DatabasePostgreSQL Database

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{
		Driver: "PostgreSQL",
	})
}

type DatabaseMySQL Database

func NewDatabaseMySQL() *DatabaseMySQL {
	return (*DatabaseMySQL)(&Database{
		Driver: "MySQL",
	})
}

type DatabaseOracle Database

func NewDatabaseOracle() *DatabaseOracle {
	return (*DatabaseOracle)(&Database{
		Driver: "Oracle",
	})
}

type UserRepository struct {
	Log bool
}

func NewUserRepository(log bool) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

type UserService struct {
	UserRepository     *UserRepository
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMySQL      *DatabaseMySQL
	DatabaseOracle     *DatabaseOracle
}

func NewUserService(userRepository *UserRepository, databasePostgreSQL *DatabasePostgreSQL, databaseMySQL *DatabaseMySQL, databaseOracle *DatabaseOracle) *UserService {
	return &UserService{
		UserRepository:     userRepository,
		DatabasePostgreSQL: databasePostgreSQL,
		DatabaseMySQL:      databaseMySQL,
		DatabaseOracle:     databaseOracle,
	}
}
