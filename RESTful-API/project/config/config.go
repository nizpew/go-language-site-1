package config

import (
    "github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
    App      AppConfig
    Database DBConfig
    JWT      JWTConfig
}

// AppConfig holds the application-specific configuration
type AppConfig struct {
    Name string
}

// DBConfig holds the database configuration
type DBConfig struct {
    Host     string
    Port     int
    Name     string
    Username string
    Password string
}

// JWTConfig holds the JWT configuration
type JWTConfig struct {
    Expired int
    Issuer  string
    Secret  string
}

// Load loads configuration from a YAML file
func Load(path string) (*Config, error) {
    var config Config

    // Set default values
    viper.SetDefault("app.name", "go-restful-api")
    viper.SetDefault("server.host", "localhost")
    viper.SetDefault("server.port", 8080)

    // Read configuration from file
    viper.SetConfigFile(path)
    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    // Unmarshal configuration
    if err := viper.Unmarshal(&config); err != nil {
        return nil, err
    }

    return &config, nil
}

