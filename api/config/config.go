package config

import "time"

type Configurations struct {
	Postgres    PostgresConfig    `mapstructure:"postgres"`
	AppConfig   AppConfig         `mapstructure:"app"`
	Payos       PayosConfig       `mapstructure:"payos"`
	SMTP        SMTPConfig        `mapstructure:"smtp"`
	Cloundinary CloundinaryConfig `mapstructure:"clouldinary"`
}

type PostgresConfig struct {
	ConnectionString string `mapstructure:"connectionString"`
}

type AppConfig struct {
	JWTSecret            string        `mapstructure:"jwtSecret"`
	JWTDuration          time.Duration `mapstructure:"jwtDuration"`
	RefreshTokenDuration time.Duration `mapstructure:"refreshTokenDuration"`
}

type PayosConfig struct {
	ClientID    string `mapstructure:"clientID"`
	APIKey      string `mapstructure:"apiKey"`
	ChecksumKey string `mapstructure:"checksumKey"`
	BaseURL     string `mapstructure:"baseURL"`
	Domain      string `mapstructure:"domain"`
}

type SMTPConfig struct {
	SMTPHost     string `mapstructure:"host"`
	SMTPPort     string `mapstructure:"port"`
	SMTPUsername string `mapstructure:"username"`
	SMTPPassword string `mapstructure:"password"`
	SenderEmail  string `mapstructure:"from"`
}

type CloundinaryConfig struct {
	Name      string `mapstructure:"name"`
	ApiKey    string `mapstructure:"apiKey"`
	ApiSecret string `mapstructure:"apiSecret"`
}
