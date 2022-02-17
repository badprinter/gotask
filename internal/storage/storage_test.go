package storage

import (
	"gotask/internal/config"
	"testing"
)

// TODO переписать, вынести поля toml в переменые.
func TestNewStorage(t *testing.T) {

	dataConf := "[db]\nip=\"localhost\"\nport=\"5432\"\nuser=\"postgresql\"\npassword=\"2132\"\ndb_name=\"gotest\""
	example := &Storage{
		Config: &config.Config{
			DB_CONFIG: &config.DB_Config{
				IP:       "localhost",
				PORT:     "5432",
				USER:     "postgresql",
				PASSWORD: "2132",
				DB_NAME:  "gotest",
			},
		},
	}

	s, _ := NewStorage(dataConf)

	if example.Config.DB_CONFIG.DB_NAME != s.Config.DB_CONFIG.DB_NAME {
		t.Errorf("exept: %s -- result: %s\n", example.Config.DB_CONFIG.DB_NAME, s.Config.DB_CONFIG.DB_NAME)
	}

}
