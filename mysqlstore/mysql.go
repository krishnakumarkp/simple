package mysqlstore

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Host, Port, User, Password, Database string
}

type DataStore struct {
	Db  *sql.DB
	cfg Config
}

func New(cfg Config) (DataStore, error) {
	var store DataStore
	if cfg.Host == "" || cfg.Port == "" || cfg.User == "" ||
		cfg.Password == "" || cfg.Database == "" {
		err := errors.New("All configuration fields must be set")
		return store, err
	}
	store.cfg = cfg

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database))

	if err != nil {
		err := errors.New("Couldnot connect to the databse")
		return store, err
	}

	// Ping verifies if the connection to the database is alive or if a
	// new connection can be made.
	if err := db.Ping(); err != nil {
		err = errors.New("Could not ping the mysql database")
		return store, err
	}

	store.Db = db
	return store, nil

}
