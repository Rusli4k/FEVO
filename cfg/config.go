// Package cfg contains struct
// that will hold on all needful parameters for our app
// that will be retrieved from .env
package cfg

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Allowed logger levels & config key.
const (
	lenOfLines  = 2
	envFileName = ".env"
)

// Load configs from a env file & sets them in environment variables.
func loadEnvVar() error {
	f, err := os.Open(envFileName)
	if err != nil {
		return fmt.Errorf("error while opening %s file: %w", envFileName, err)
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Printf("%s", err)
		}
	}()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error while scanning %s file: %w", envFileName, err)
	}

	for _, l := range lines {
		pair := strings.Split(l, "=")
		if len(pair) != lenOfLines {
			return fmt.Errorf("not enough data for the configuration at the config file")
		}

		os.Setenv(pair[0], pair[1])
	}

	return nil
}

// DB describes DB parameters.
type DB struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

// Server configuration description.
type Server struct {
	Host string
	Port string
}

// Options will keep all needful configs.
type Options struct {
	Server Server
	DB     DB
}

// GetConfig will create instance of Options
// that will be used im main package.
func GetConfig() (Options, error) {
	if err := loadEnvVar(); err != nil {
		return Options{}, err
	}

	opt := Options{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		DB: DB{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
	}

	return opt, nil
}
