package model

type SignUp struct {
	Login            string
	Password         string
	VerificationCode string
}

type SignIn struct {
	Login    string
	Password string
}

type SignOut struct {
	RefreshToken string
}

type Jwt struct {
	AccessToken  string
	RefreshToken string
}
