package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DBs struct {
	MySQL *sql.DB
}

func NewDBs(cfg map[string]struct {
	Driver string
	DSN    string
}) (*DBs, error) {
	dbs := &DBs{}
	for name, c := range cfg {
		db, err := sql.Open(c.Driver, c.DSN)
		if err != nil {
			return nil, fmt.Errorf("open %s DB error: %w", name, err)
		}
		if err := db.Ping(); err != nil {
			return nil, fmt.Errorf("ping %s DB error: %w", name, err)
		}

		if name == "mysql" {
			dbs.MySQL = db
		}
	}
	return dbs, nil
}
