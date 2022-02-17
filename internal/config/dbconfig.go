package config

import "fmt"

type DB_Config struct {
	IP       string
	PORT     string
	USER     string
	PASSWORD string
	DB_NAME  string
	SSLMODE  string
}

func (db *DB_Config) GetUrlPostgres() string {
	return fmt.Sprintf("user=%s password='%s' dbname=%s sslmode=%s", db.USER, db.PASSWORD, db.DB_NAME, db.SSLMODE)
}
