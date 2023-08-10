package config

import (
	"fmt"

	"github.com/caarlos0/env/v9"
)

type Database struct {
	username string `env:"DATABASE_USERNAME"`
}

func (Database *Database) UseConfig() {
	err := env.Parse(Database)
	if err != nil {
		fmt.Println(err)
	}
}
