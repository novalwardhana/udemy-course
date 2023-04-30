//go:build wireinject

// + wireinject

package simple

import "github.com/google/wire"

func InitializeCategoryRepository() *CategoryRepository {
	wire.Build(
		NewCategoryService,
		wire.FieldsOf(new(*CategoryService), "CategoryRepository"),
	)
	return nil
}

func mu() *ManchesterUnited {
	wire.Build(
		NewClub,
		wire.FieldsOf(new(*Club), "ManchesterUnited"),
	)
	return nil
}

func chelsea() *Chelsea {
	wire.Build(
		NewClub,
		wire.FieldsOf(new(*Club), "Chelsea"),
	)
	return nil
}

func arsenal() *Arsenal {
	wire.Build(
		NewClub,
		wire.FieldsOf(new(*Club), "Arsenal"),
	)
	return nil
}
