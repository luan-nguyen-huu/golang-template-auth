package configs

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:",squash"`
	DB       DBConfig       `mapstructure:",squash"`
	JWT      JWTConfig      `mapstructure:",squash"`
	Password PasswordConfig `mapstructure:",squash"`
	CORS     CORSConfig     `mapstructure:",squash"`
}

type AppConfig struct {
	Name string `mapstructure:"APP_NAME"`
	Env  string `mapstructure:"APP_ENV"`
	Host string `mapstructure:"APP_HOST"`
	Port int    `mapstructure:"APP_PORT"`
}

type DBConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
	SSLMode  string `mapstructure:"DB_SSLMODE"`
}

type JWTConfig struct {
	SecretAccess       string        `mapstructure:"JWT_SECRET_ACCESS"`
	SecretRefresh      string        `mapstructure:"JWT_SECRET_REFRESH"`
	AccessTokenExpire  time.Duration `mapstructure:"JWT_ACCESS_TOKEN_EXPIRE"`
	RefreshTokenExpire time.Duration `mapstructure:"JWT_REFRESH_TOKEN_EXPIRE"`
	Secure             bool          `mapstructure:"JWT_COOKIE_SECURE"`
}

type PasswordConfig struct {
	HashCost int `mapstructure:"PASSWORD_HASH_COST"`
}

type CORSConfig struct {
	AllowedOrigins string `mapstructure:"CORS_ALLOWED_ORIGINS"`
}

// GetAllowedOrigins returns slice of allowed CORS origins.
func (c *CORSConfig) GetAllowedOrigins() []string {
	if c.AllowedOrigins == "" {
		return []string{"http://localhost:3000", "http://localhost:8080"}
	}
	origins := strings.Split(c.AllowedOrigins, ",")
	for i := range origins {
		origins[i] = strings.TrimSpace(origins[i])
	}
	return origins
}

func Load() (*Config, error) {
	v := viper.New()
	v.SetConfigFile(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Default fallback values
	v.SetDefault("APP_PORT", 8080)
	v.SetDefault("APP_HOST", "0.0.0.0")
	v.SetDefault("APP_ENV", "development")
	v.SetDefault("DB_SSLMODE", "disable")
	v.SetDefault("JWT_ACCESS_TOKEN_EXPIRE", "15m")
	v.SetDefault("JWT_REFRESH_TOKEN_EXPIRE", "24h")
	v.SetDefault("JWT_COOKIE_SECURE", false)
	v.SetDefault("PASSWORD_HASH_COST", 10)
	v.SetDefault("CORS_ALLOWED_ORIGINS", "http://localhost:3000,http://localhost:8080")

	_ = v.ReadInConfig() // If .env doesn't exist, read from system env

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return &cfg, nil
}

func GetPostgresDSN(cfg *Config) string {
	sslMode := cfg.DB.SSLMode
	if sslMode == "" {
		sslMode = "disable"
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		sslMode,
	)
}