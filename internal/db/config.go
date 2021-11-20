package db

import "fmt"

type ConfigDB struct {
	Host     string
	User     string
	Password string
	Name     string
}

func (cfg *ConfigDB) String() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.Name,
	)
}
