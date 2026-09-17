package repository

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r UserRepository) Get() {
}

func (r UserRepository) GetMany() {
}

func (r UserRepository) Create() {
}

func (r UserRepository) Update() {
}

func (r UserRepository) Delete() {
}
