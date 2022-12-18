package simple

type FooRepository struct {
}

func NewFooRepository() *FooRepository {
	return &FooRepository{}
}

type FooService struct {
	*FooRepository
}

func NewFooService(foorepository *FooRepository) *FooService {
	return &FooService{foorepository}
}
