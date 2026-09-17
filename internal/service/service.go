package service

type Service struct {
	Auth *AuthService
	User *UserService
}

func NewService() *Service {
	return &Service{
		Auth: NewAuthService(),
		User: NewUserService(),
	}
}
