package auth_token

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"social/internal/pkg/apperror"
)

type service struct {
	rep   Repository
}

type Service interface {
	Get(ctx context.Context, token string) (c *AuthToken, err error)
	Generate(ctx context.Context, userID uint) (c *AuthToken, err error)
}

func NewService(rep Repository) Service {
	return &service{
		rep: rep,
	}
}

func (m service) Get(ctx context.Context, token string) (authToken *AuthToken, err error) {
	if authToken, err = m.rep.Get(ctx, token); err != nil {
		return nil, err
	}

	tm := time.Unix(int64(authToken.ExpiredAt), 0)
	if tm.Before(time.Now()) {
		return nil, apperror.ErrTokenExpired
	}

	return authToken, nil
}

func (m service) Generate(ctx context.Context, userID uint) (c *AuthToken, err error) {
	hasher := sha1.New()
	hasher.Write([]byte(fmt.Sprintf("%d%d", time.Now().Unix(), userID)))
	hash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	authToken := AuthToken{
		Token:     hash,
		UserId:    userID,
		ExpiredAt: uint(time.Now().Unix()) + timeExpiredSeconds,
	}

	err = m.rep.Create(ctx, &authToken)
	if err != nil {
		return nil, errors.Wrap(err, "cannot save auth token")
	}

	return &authToken, nil
}