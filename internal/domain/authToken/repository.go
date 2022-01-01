package user

import (
	"context"

	minipkg_gorm "github.com/minipkg/db/gorm"
	"github.com/minipkg/selection_condition"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"social/internal/pkg/apperror"
)

type Repository interface {
	Get(ctx context.Context, token string) (c *AuthToken, err error)
	Generate(ctx context.Context, cond *AuthToken) (c *AuthToken, err error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (m repository) Get(ctx context.Context, id uint) (b *AuthToken, err error) {
	b = &AuthToken{}
	err = m.db.WithContext(ctx).First(b, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}
		return nil, errors.Wrap(err, "cannot get by id AuthToken")
	}
	return b, nil
}

func (m repository) First(ctx context.Context, cond *AuthToken) (b *AuthToken, err error) {
	b = &AuthToken{}
	err = m.db.WithContext(ctx).Where(cond).First(b).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}
		return nil, errors.Wrap(err, "cannot get AuthToken")
	}
	return b, nil
}

func (m repository) Query(ctx context.Context, cond *selection_condition.SelectionCondition) (AuthTokens []AuthToken, err error) {
	db := m.db.WithContext(ctx)
	db = minipkg_gorm.Conditions(db, cond)

	err = db.Find(&AuthTokens).Error
	if err != nil {
		return nil, errors.Wrap(err, "cannot get AuthTokens")
	}
	return AuthTokens, nil
}

func (m repository) Create(ctx context.Context, c *AuthToken) (uint, error) {
	err := m.db.WithContext(ctx).Create(c).Error
	if err != nil {
		return 0, errors.Wrap(err, "cannot create AuthToken")
	}
	return c.ID, nil
}

func (m repository) Update(ctx context.Context, c *AuthToken) error {
	err := m.db.WithContext(ctx).Save(c).Error
	if err != nil {
		return errors.Wrap(err, "cannot update AuthToken")
	}
	return nil
}

func (m repository) Delete(ctx context.Context, id uint) error {
	b, err := m.Get(ctx, id)
	if err != nil {
		return err
	}

	err = m.db.WithContext(ctx).Model(b).Select(clause.Associations).Delete(b).Error
	if err != nil {
		return errors.Wrap(err, "cannot delete AuthToken")
	}
	return nil
}
