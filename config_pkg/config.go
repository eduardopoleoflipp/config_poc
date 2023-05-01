package config_pkg

/*
	Copied from
	https://github.com/wishabi/pkg/blob/main/config/loader.go
	for convience to show case the POC
*/

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Environment string

const (
	Production Environment = "production"
	Docker     Environment = "docker"
	Test       Environment = "test"
	Staging    Environment = "staging"
	Local      Environment = "local"
)

// GetEnvironment returns the current environment variable
func GetEnvironment() Environment {
	switch strings.ToLower(os.Getenv("ENVIRONMENT")) {
	case "production":
		return Production
	case "docker":
		return Docker
	case "test":
		return Test
	case "staging":
		return Staging
	default:
		return Local
	}
}

// Loader accepts a config path and tries to load all configuration inside that path
// any field inside the configuration can be updated using FLIPP_<name of config field> environment
// variables
func Loader(configPath string, ptr interface{}) error {
	var err error

	viper := viper.New()

	var filePattern string

	switch GetEnvironment() {
	case Production:
		filePattern = "config.prod"
	case Docker:
		filePattern = "config.docker"
	case Test:
		filePattern = "config.test"
	case Staging:
		filePattern = "config.staging"
	default:
		filePattern = "config.local"
	}

	viper.SetConfigName(filePattern)

	log.
		Debug().
		Str("environment", os.Getenv("ENVIRONMENT")).
		Str("file_pattern", filePattern).
		Msg("loading configuration file based on env variable")

	viper.SetEnvPrefix("FLIPP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if configPath == "" {
		err = viper.ReadConfig(strings.NewReader("{}"))
	} else {
		viper.AddConfigPath(configPath)
		err = viper.ReadInConfig()
	}

	if err != nil {
		return err
	}

	err = viper.Unmarshal(ptr)
	if err != nil {
		return err
	}

	return nil
}
