package config

import "time"

type Configurations struct {
	Postgres  PostgresConfig `mapstructure:"postgres"`
	AppConfig AppConfig      `mapstructure:"app"`
}

type PostgresConfig struct {
	ConnectionString string `mapstructure:"connectionString"`
}

type AppConfig struct {
	JWTSecret   string        `mapstructure:"jwtSecret"`
	JWTDuration time.Duration `mapstructure:"jwtDuration"`
}
