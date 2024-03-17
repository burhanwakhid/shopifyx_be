package orm

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	*sql.DB
}

type Option struct {
	Host         string
	Port         int
	User         string
	Password     string
	Database     string
	MaxLifeTime  time.Duration
	MaxIdleConns int
	MaxOpenConns int
	LogLevel     int
}

func (o Option) validate() error {
	if o.Host == "" {
		return errors.New("host is required")
	}
	if o.Port == 0 {
		return errors.New("port is required")
	}
	if o.User == "" {
		return errors.New("user is required")
	}
	if o.Database == "" {
		return errors.New("database is required")
	}
	return nil
}

func (o Option) connectionString() string {

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		o.Host, o.Port, o.User, o.Password, o.Database)

}

func NewDB(o Option) (*DB, error) {
	err := o.validate()
	if err != nil {
		return nil, err
	}

	conn, err := sql.Open("postgres", o.connectionString())
	if err != nil {
		return nil, err
	}

	conn.SetConnMaxLifetime(o.MaxLifeTime)
	conn.SetMaxIdleConns(o.MaxIdleConns)
	conn.SetMaxOpenConns(o.MaxOpenConns)

	return &DB{conn}, nil
}

func Open(o Option) (*DB, error) {
	// err := o.validate()
	// if err != nil {
	// 	return nil, err
	// }

	conn, err := sql.Open("postgres", o.Database)
	if err != nil {
		return nil, err
	}

	conn.SetConnMaxLifetime(o.MaxLifeTime)
	conn.SetMaxIdleConns(o.MaxIdleConns)
	conn.SetMaxOpenConns(o.MaxOpenConns)

	return &DB{conn}, nil
}
