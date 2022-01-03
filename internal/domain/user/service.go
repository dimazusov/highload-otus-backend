package user

import (
	"context"
	"github.com/minipkg/selection_condition"
	"social/internal/pkg/apperror"
	"social/internal/pkg/hasher"
	"social/internal/pkg/jwt_token"
)

type Authentifier interface {
	Check(token string) (bool, error)
	Login(credentials Credentials) (token uint, err error)
	Logout(token string) error
	GetUserID(token string) (userID uint, err error)
}

type service struct {
	rep          Repository
	authentifier Authentifier
}

type Service interface {
	Get(ctx context.Context, id uint) (b *User, err error)
	First(ctx context.Context, cond *User) (b *User, err error)
	Query(ctx context.Context, cond *selection_condition.SelectionCondition) (Users []User, err error)
	Create(ctx context.Context, b *User) (uint, error)
	Update(ctx context.Context, b *User) error
	Delete(ctx context.Context, id uint) error
	Count(ctx context.Context, cond *selection_condition.SelectionCondition) (uint, error)
	Auth(ctx context.Context, credentials Credentials) (token string, err error)
	Logout(ctx context.Context, token string) (err error)
}

func NewService(rep Repository) Service {
	return &service{rep: rep}
}

func (m service) Get(ctx context.Context, id uint) (b *User, err error) {
	if b, err = m.rep.Get(ctx, id); err != nil {
		return nil, err
	}

	return b, nil
}

func (m service) First(ctx context.Context, cond *User) (b *User, err error) {
	if b, err = m.rep.Get(ctx, cond.ID); err != nil {
		return nil, err
	}

	return b, nil
}

func (m service) Query(ctx context.Context, cond *selection_condition.SelectionCondition) (Users []User, err error) {
	if Users, err = m.rep.Query(ctx, cond); err != nil {
		return nil, err
	}

	return Users, nil
}

func (m service) Create(ctx context.Context, b *User) (uint, error) {
	return m.rep.Create(ctx, b)
}

func (m service) Update(ctx context.Context, b *User) error {
	return m.rep.Update(ctx, b)
}

func (m service) Delete(ctx context.Context, id uint) error {
	return m.rep.Delete(ctx, id)
}

func (m service) Count(ctx context.Context, cond *selection_condition.SelectionCondition) (uint, error) {
	return m.rep.Count(ctx, cond)
}

func (m service) Auth(ctx context.Context, credentials Credentials) (token string, err error) {
	passHash, err := hasher.GetHashFromStruct(credentials.Password)
	if err != nil {
		return "", err
	}

	user, err := m.rep.First(ctx, &User{Email: credentials.Email, Password: passHash})
	if err != nil {
		return "", apperror.ErrUserNotFound
	}

	p := jwt_token.NewParser()
	jwtToken, err := p.Create(user.ID)
	if err != nil {
		return "", err
	}

	token, err = jwtToken.ToString()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (m service) Logout(ctx context.Context, token string) error {
	return nil
}
