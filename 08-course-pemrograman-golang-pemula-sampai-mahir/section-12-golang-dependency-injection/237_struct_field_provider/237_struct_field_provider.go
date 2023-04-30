package simple

/* Example 1 */
type CategoryRepository struct {
	Name string
}

type CategoryService struct {
	CategoryRepository *CategoryRepository
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		CategoryRepository: &CategoryRepository{
			Name: "Noval Wardhana",
		},
	}
}

/* Example 2 */
type ManchesterUnited struct {
	Name string
}

type Chelsea struct {
	Name string
}

type Arsenal struct {
	Name string
}

type Club struct {
	*ManchesterUnited
	*Chelsea
	*Arsenal
}

func NewClub() *Club {
	return &Club{
		ManchesterUnited: &ManchesterUnited{},
		Chelsea:          &Chelsea{},
		Arsenal:          &Arsenal{},
	}
}
