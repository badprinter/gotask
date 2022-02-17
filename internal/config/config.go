package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	DB_CONFIG *DB_Config `toml:"db"`
}

func NewCofnig(RawConf string) (*Config, error) {
	var conf Config
	if _, err := toml.Decode(RawConf, &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

func (c *Config) String() string {
	return fmt.Sprintf("Data Base config:\n\tIP - %s\n\tPORT - %s\n\tUSER - %s\n\tPASSWORD - %s\n\tDB_NAME - %s\n\tSSLMODE - %s", c.DB_CONFIG.IP, c.DB_CONFIG.PORT, c.DB_CONFIG.USER, c.DB_CONFIG.PASSWORD, c.DB_CONFIG.DB_NAME, c.DB_CONFIG.SSLMODE)
}
