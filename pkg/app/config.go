package app

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
)

// Config stores the application-wide configurations
var Config appConfig

type appConfig struct {
	// the path to the error message file. Defaults to "config/errors.yaml"
	ErrorFile string `mapstructure:"error_file"`
	// running mode
	MODE string `mapstructure:"mode"`
	// the server port. Defaults to 8080
	ServerPort int `mapstructure:"server_port"`
	// the host of db
	DBHost string `mapstructure:"db_host"`
	// the port of db
	DBPort string `mapstructure:"db_port"`
	// the name of db
	DBName string `mapstructure:"db_name"`
	// the username of db
	DBUser string `mapstructure:"db_user"`
	// the password of db
	DBPass string `mapstructure:"db_pass"`
}

func (config appConfig) Validate() error {
	return validation.ValidateStruct(&config,
		validation.Field(&config.DBName, validation.Required),
		validation.Field(&config.DBUser, validation.Required),
		validation.Field(&config.DBPass, validation.Required),
	)
}

// LoadConfig loads configuration from the given list of paths and populates it into the Config variable.
// The configuration file(s) should be named as app.yaml.
// Environment variables with the prefix "RESTFUL_" in their names are also read automatically.
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("mc")
	v.AutomaticEnv()

	// set default values
	v.SetDefault("mode", "debug")
	v.SetDefault("error_file", "config/errors.yaml")
	v.SetDefault("server_port", 8080)
	v.SetDefault("db_host", "127.0.0.1")
	v.SetDefault("db_port", "3306")

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}

	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to read the configuration file: %s", err)
	}

	if err := v.Unmarshal(&Config); err != nil {
		return err
	}

	return Config.Validate()
}
