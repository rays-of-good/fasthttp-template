package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type (
	Configuration struct {
		DSN string
	}

	Database struct {
		Pool *pgxpool.Pool
	}
)

func NewDatabase(c *Configuration) (d *Database, err error) {
	p, err := pgxpool.Connect(context.Background(), c.DSN)
	if err != nil {
		return
	}

	d = &Database{
		Pool: p,
	}

	return
}

func (d *Database) Close() (err error) {
	d.Pool.Close()
	return
}
