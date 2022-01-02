package auth_token

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"social/internal/pkg/apperror"
)

type Repository interface {
	Get(ctx context.Context, token string) (c *AuthToken, err error)
	Create(ctx context.Context, authToken *AuthToken) (err error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (m repository) Get(ctx context.Context, token string) (c *AuthToken, err error) {
	authToken := &AuthToken{Token: token}
	err = m.db.WithContext(ctx).First(authToken, authToken).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}
		return nil, errors.Wrap(err, "cannot get by id auth token")
	}

	return authToken, nil
}

func (m repository) Create(ctx context.Context, authToken *AuthToken) (err error) {
	err = m.db.WithContext(ctx).Save(authToken).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.ErrNotFound
		}
		return errors.Wrap(err, "cannot get by id auth token")
	}

	return  nil
}
