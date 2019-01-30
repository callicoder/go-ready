package config

import (
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Logging  LoggingConfig
	Database DatabaseConfig
	Auth     AuthConfig
	Grpc     GrpcConfig
}

type ServerConfig struct {
	ContextPath                string
	Port                       int
	ReadTimeoutSec             int
	WriteTimeoutSec            int
	GracefulShutdownTimeoutSec int
}

type GrpcConfig struct {
	Port                       int
	ConnectionTimeoutSec       int
	GracefulShutdownTimeoutSec int
}

type LoggingConfig struct {
	Level string
}

type DatabaseConfig struct {
	DriverName string
	URL        string
}

type AuthConfig struct {
	JwtSecret      string
	JwtExpiryInSec int
}

func Load(configFile string) (*Config, error) {
	fileName := filepath.Base(configFile)
	dir := filepath.Dir(configFile)
	ext := filepath.Ext(configFile)
	fileNameWithoutExt := fileName[0 : len(fileName)-len(ext)]

	viper.AutomaticEnv()
	viper.SetConfigName(fileNameWithoutExt)

	if ext != "" {
		viper.SetConfigType(strings.TrimLeft(ext, "."))
	} else {
		viper.SetConfigType("yml")
	}

	if dir != "" {
		viper.AddConfigPath(dir)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("config")
	}

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	var config Config
	unmarshalErr := viper.Unmarshal(&config)

	return &config, unmarshalErr
}
