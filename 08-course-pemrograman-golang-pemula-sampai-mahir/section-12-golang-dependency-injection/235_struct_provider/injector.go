//go:build wireinject

// + wireinject

package simple

import "github.com/google/wire"

func InitializedCategoryClass() *CategoryClass {
	wire.Build(NewCategory, NewClass,
		wire.Struct(new(CategoryClass), "*"),
	)
	return nil
}

func InitializedDatabase() *DatabaseRepository {
	wire.Build(NewDatabasePostgreSQL, NewDatabaseMySQL, NewDatabaseOracle,
		wire.Struct(new(DatabaseRepository), "pg", "mysql", "oracle"),
	)
	return nil
}
