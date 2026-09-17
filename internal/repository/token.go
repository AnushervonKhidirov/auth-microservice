package repository

type TokenRepository struct{}

func NewTokenRepository() *TokenRepository {
	return &TokenRepository{}
}

func (r TokenRepository) Get() {
}

func (r TokenRepository) Create() {
}

func (r TokenRepository) Delete() {
}

func (r TokenRepository) DeleteExpired() {
}
