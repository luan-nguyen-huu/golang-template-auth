package configs

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
)

var (
	Cfg, _    = Load()
)

type Config struct {
	App struct {
		Name string `mapstructure:"APP_NAME"`
		Env  string `mapstructure:"APP_ENV"`
		Host string `mapstructure:"APP_HOST"`
		Port int    `mapstructure:"APP_PORT"`
	} `mapstructure:",squash"`

	DB struct {
		Host     string `mapstructure:"DB_HOST"`
		Port     int    `mapstructure:"DB_PORT"`
		User     string `mapstructure:"DB_USER"`
		Password string `mapstructure:"DB_PASSWORD"`
		Name     string `mapstructure:"DB_NAME"`
	} `mapstructure:",squash"`

	JWT struct {
		SecretAccess           string `mapstructure:"JWT_SECRET_ACCESS"`
		SecretRefresh          string `mapstructure:"JWT_SECRET_REFRESH"`
		AccessTokenExpire  time.Duration    `mapstructure:"JWT_ACCESS_TOKEN_EXPIRE"`
		RefreshTokenExpire time.Duration    `mapstructure:"JWT_REFRESH_TOKEN_EXPIRE"`
		Secure     bool          `env:"SECURE, default=true"`
	} `mapstructure:",squash"`

	Password struct {
		HashCost int `mapstructure:"PASSWORD_HASH_COST"`
	} `mapstructure:",squash"`
}

func Load() (*Config, error) {
	v := viper.New()
	v.SetConfigFile(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse env: %w", err)
	}
	return &cfg, nil
}


func GetPostgresDSN(cfg *Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
	)
}