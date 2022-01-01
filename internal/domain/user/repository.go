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
	Get(ctx context.Context, id uint) (c *User, err error)
	First(ctx context.Context, cond *User) (c *User, err error)
	Query(ctx context.Context, cond *selection_condition.SelectionCondition) (Users []User, err error)
	Create(ctx context.Context, c *User) (uint, error)
	Update(ctx context.Context, c *User) error
	Delete(ctx context.Context, id uint) error
	Count(ctx context.Context, cond *selection_condition.SelectionCondition) (uint, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (m repository) Get(ctx context.Context, id uint) (b *User, err error) {
	b = &User{}
	err = m.db.WithContext(ctx).First(b, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}
		return nil, errors.Wrap(err, "cannot get by id User")
	}
	return b, nil
}

func (m repository) First(ctx context.Context, cond *User) (b *User, err error) {
	b = &User{}
	err = m.db.WithContext(ctx).Where(cond).First(b).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrNotFound
		}
		return nil, errors.Wrap(err, "cannot get User")
	}
	return b, nil
}

func (m repository) Query(ctx context.Context, cond *selection_condition.SelectionCondition) (Users []User, err error) {
	db := m.db.WithContext(ctx)
	db = minipkg_gorm.Conditions(db, cond)

	err = db.Find(&Users).Error
	if err != nil {
		return nil, errors.Wrap(err, "cannot get Users")
	}
	return Users, nil
}

func (m repository) Create(ctx context.Context, c *User) (uint, error) {
	err := m.db.WithContext(ctx).Create(c).Error
	if err != nil {
		return 0, errors.Wrap(err, "cannot create User")
	}
	return c.ID, nil
}

func (m repository) Update(ctx context.Context, c *User) error {
	err := m.db.WithContext(ctx).Save(c).Error
	if err != nil {
		return errors.Wrap(err, "cannot update User")
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
		return errors.Wrap(err, "cannot delete User")
	}
	return nil
}

func (m repository) Count(ctx context.Context, cond *selection_condition.SelectionCondition) (uint, error) {
	db := minipkg_gorm.Conditions(m.db, cond)

	var count int64
	err := db.WithContext(ctx).Model(&User{}).Count(&count).Error
	if err != nil {
		return 0, errors.Wrap(err, "cannot get count Users")
	}

	return uint(count), nil
}
