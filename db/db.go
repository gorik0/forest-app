package db

import (
	"awesomeProject/config"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectDb(cfg *config.Config) (*sql.DB, error) {

	dsn := fmt.Sprintf("host=%s  dbname=%s password=%s  user=%s sslmode=disable", cfg.DbHost, cfg.DBName, cfg.DbPassword, cfg.DbUser)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, errors.New("Fail to connect db --->>>>" + err.Error())
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.New("Fail to ping db --->>>>" + err.Error())
	}

	return db, nil
}
