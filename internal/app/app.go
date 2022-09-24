package app

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"social/internal/cache"
	"social/internal/config"
	"social/internal/domain/auth_token"
	"social/internal/domain/user"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type Domain struct {
	User      DomainUser
	AuthToken DomainAuthToken
}

type DomainAuthToken struct {
	Service auth_token.Service
}

type DomainUser struct {
	Repository user.Repository
	Service    user.Service
}

type App struct {
	cfg    *config.Config
	db     *sql.DB
	Cache  cache.Cache
	Domain Domain
}

func New(config *config.Config) *App {
	return &App{cfg: config}
}

func (m *App) DB() *sql.DB {
	return m.db
}

func (m *App) Init() error {
	if err := m.initDB(); err != nil {
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

func (m *App) initDB() error {
	switch m.cfg.Repository.Type {
	case "mysql":
		db, err := sql.Open(m.cfg.DB.Mysql.Dialect, m.cfg.DB.Mysql.Dsn)
		if err != nil {
			return errors.Wrap(err, "cannot connect to db")
		}
		db.SetConnMaxLifetime(time.Minute * 3)
		db.SetMaxOpenConns(m.cfg.DB.Mysql.MaxConn)
		db.SetMaxIdleConns(m.cfg.DB.Mysql.MaxConn)
		m.db = db
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

func (m *App) initRepositories() (err error) {
	m.Domain.User.Repository = user.NewRepository(m.db)

	return nil
}

func (m *App) initServices() (err error) {
	m.Domain.User.Service = user.NewService(m.Domain.User.Repository)

	return nil
}
