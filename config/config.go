package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	AppName string `mapstructure:"app_name"`
	Port    string `mapstructure:"port"`
	DB      DBConfig
}

// DBConfig holds the database-related configuration
type DBConfig struct {
	Host     string `mapstructure:"db_host"`
	Port     string `mapstructure:"db_port"`
	User     string `mapstructure:"db_user"`
	Password string `mapstructure:"db_password"`
	Name     string `mapstructure:"db_name"`
}

// LoadConfig reads configuration from file and environment variables
func LoadConfig(configPath string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("././")

	viper.SetEnvPrefix("MYAPP")
	viper.AutomaticEnv()

	// Set default values
	viper.SetDefault("app_name", "MyApp")
	viper.SetDefault("port", 8080)
	viper.SetDefault("db_host", "localhost")
	viper.SetDefault("db_port", 5432)
	viper.SetDefault("db_user", "user")
	viper.SetDefault("db_password", "password")
	viper.SetDefault("db_name", "mydatabase")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: No config file found, using defaults & environment variables.")
	}
	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}

	return &config, nil

}
