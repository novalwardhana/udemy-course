package simple

/* Example 1 */
type Category struct {
}

func NewCategory() *Category {
	return &Category{}
}

type Class struct {
}

func NewClass() *Class {
	return &Class{}
}

type CategoryClass struct {
	*Category
	*Class
}

/* Example 2 */
type Database struct {
	Driver string
	Name   string
}

type DatabasePostgreSQL Database

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return &DatabasePostgreSQL{
		Driver: "pg",
	}
}

type DatabaseMySQL Database

func NewDatabaseMySQL() *DatabaseMySQL {
	return &DatabaseMySQL{
		Driver: "mysql",
	}
}

type DatabaseOracle Database

func NewDatabaseOracle() *DatabaseOracle {
	return &DatabaseOracle{
		Driver: "oci",
	}
}

type DatabaseRepository struct {
	pg     *DatabasePostgreSQL
	mysql  *DatabaseMySQL
	oracle *DatabaseOracle
}
