package camagru

import (
	"errors"
	"log"
	"os"
	"strconv"
)

// SMTPConfig ...
type SMTPConfig struct {
	From     string
	Password string
	Host     string
	Port     int
}

// DatabaseConfig ...
type DatabaseConfig struct {
	Host    string
	Port    int
	User    string
	Name    string
	SSLMode string
}

// Config ...
type Config struct {
	BindAddr   string
	SessionKey string
	Database   *DatabaseConfig
	SMTP       *SMTPConfig
}

// NewConfig ...
func NewConfig() (*Config, error) {
	bindAddr, exists := os.LookupEnv("BIND_ADDR")
	if !exists {
		return nil, errors.New("BIND_ADDR not found")
	}

	sessionKey, exists := os.LookupEnv("SESSION_KEY")
	if !exists {
		return nil, errors.New("SESSION_KEY not found")
	}

	smtpConfig, err := NewSMTPConfig()
	if err != nil {
		return nil, err
	}

	databaseConfig, err := NewDatabaseConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		BindAddr:   bindAddr,
		Database:   databaseConfig,
		SMTP:       smtpConfig,
		SessionKey: sessionKey,
	}, nil
}

// NewSMTPConfig ...
func NewSMTPConfig() (*SMTPConfig, error) {
	from, exists := os.LookupEnv("FROM")
	if !exists {
		return nil, errors.New("FROM not found")
	}

	password, exists := os.LookupEnv("PASSWORD")
	if !exists {
		return nil, errors.New("PASSWORD not found")
	}

	host, exists := os.LookupEnv("HOST")
	if !exists {
		return nil, errors.New("HOST not found")
	}

	portS, exists := os.LookupEnv("PORT")
	if !exists {
		return nil, errors.New("PORT not found")
	}

	port, err := strconv.Atoi(portS)
	if err != nil {
		log.Println(portS)
		return nil, errors.New("PORT is invalid")
	}

	return &SMTPConfig{
		From:     from,
		Password: password,
		Host:     host,
		Port:     port,
	}, nil
}

// NewDatabaseConfig ...
func NewDatabaseConfig() (*DatabaseConfig, error) {
	host, exists := os.LookupEnv("DB_HOST")
	if !exists {
		return nil, errors.New("DB_HOST not found")
	}

	portS, exists := os.LookupEnv("DB_PORT")
	if !exists {
		return nil, errors.New("DB_PORT not found")
	}

	port, err := strconv.Atoi(portS)
	if err != nil {
		return nil, errors.New("PORT is invalid")
	}

	user, exists := os.LookupEnv("DB_USER")
	if !exists {
		return nil, errors.New("HOST not found")
	}

	name, exists := os.LookupEnv("DB_NAME")
	if !exists {
		return nil, errors.New("PORT not found")
	}

	sslMode, exists := os.LookupEnv("DB_SSLMODE")
	if !exists {
		return nil, errors.New("PORT not found")
	}

	return &DatabaseConfig{
		Host:    host,
		Port:    port,
		User:    user,
		Name:    name,
		SSLMode: sslMode,
	}, nil
}
