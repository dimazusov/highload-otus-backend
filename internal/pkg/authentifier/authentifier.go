package authentifier

import "social/internal/domain/user"

type authentifier struct {}

func New() *authentifier {
	return &authentifier{}
}

func (m authentifier) Check(token string) (bool, error) {

	return
}

func (m authentifier) Login(credentials user.Credentials) (token uint, err error) {
	return
}

func (m authentifier) Logout(token string) error {
	return
}

func (m authentifier) GetUserID(token string) (userID uint, err error) {
	return
}