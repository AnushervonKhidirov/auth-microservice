package repository

type Repository struct {
	User  *UserRepository
	Token *TokenRepository
}

func NewRepository() *Repository {
	return &Repository{
		User:  NewUserRepository(),
		Token: NewTokenRepository(),
	}
}
