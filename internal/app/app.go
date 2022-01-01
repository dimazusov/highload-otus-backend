package app

import (
	"context"

	"github.com/go-redis/redis/v8"
	minipkg_gorm "github.com/minipkg/db/gorm"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"social/internal/cache"
	"social/internal/config"
	"social/internal/domain/authToken"
	"social/internal/domain/user"
)

type Domain struct {
	User      DomainUser
	AuthToken DomainAuthToken
}

type DomainAuthToken struct {
	Repository authToken.Repository
	Service    authToken.Service
}

type DomainUser struct {
	Repository user.Repository
	Service    user.Service
}

type App struct {
	cfg    *config.Config
	db     *minipkg_gorm.DB
	Cache  cache.Cache
	Domain Domain
}

func New(config *config.Config) *App {
	return &App{cfg: config}
}

func (m *App) DB() *gorm.DB {
	return m.db.DB()
}

func (m *App) Init() error {
	if err := m.initDB(); err != nil {
		return err
	}
	if err := m.initSchemes(); err != nil {
		return err
	}
	if err := m.initCache(); err != nil {
		return err
	}
	if err := m.initRepositories(); err != nil {
		return err
	}
	if err := m.initServices(); err != nil {
		return err
	}

	return nil
}

func (m *App) initDB() (err error) {
	switch m.cfg.Repository.Type {
	case "mysql":
		conn, err := gorm.Open(postgres.Open(m.cfg.DB.Postgres.Dsn), &gorm.Config{})
		if err != nil {
			return errors.Wrapf(err, "cannot connect to postgres")
		}
		m.db = &minipkg_gorm.DB{GormDB: conn}
	}

	return nil
}

func (m *App) initCache() (err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     m.cfg.Redis.Address,
		Password: m.cfg.Redis.Password,
	})
	err = rdb.Ping(context.Background()).Err()
	if err != nil {
		return errors.New("cannot connect to redis")
	}

	m.Cache = cache.New(rdb)

	return nil
}

func (m *App) initSchemes() (err error) {
	if m.db, err = m.db.SchemeInit(&user.User{}); err != nil {
		return err
	}

	return nil
}

func (m *App) initRepositories() (err error) {
	m.Domain.User.Repository = user.NewRepository(m.db.DB())

	return nil
}

func (m *App) initServices() (err error) {
	m.Domain.User.Service = user.NewService(m.Domain.User.Repository)

	return nil
}
