package config

import (
	"bytes"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server    ServerConfig
	Logging   LoggingConfig
	Database  DatabaseConfig
	Auth      AuthConfig
	Grpc      GrpcConfig
	Migration MigrationConfig
	Http      HttpConfig
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
	Driver   string
	Name     string
	Host     string
	Port     int
	Username string
	Password string
}

func (cfg DatabaseConfig) URL() string {
	var buf bytes.Buffer

	// [username[:password]@]
	if len(cfg.Username) > 0 {
		buf.WriteString(cfg.Username)
		if len(cfg.Password) > 0 {
			buf.WriteByte(':')
			buf.WriteString(cfg.Password)
		}
		buf.WriteByte('@')
	}

	// [protocol[(address)]]
	if len(cfg.Host) > 0 {
		buf.WriteString("tcp")
		buf.WriteByte('(')
		buf.WriteString(cfg.Host)
		if cfg.Port > 0 {
			buf.WriteByte(':')
			buf.WriteString(strconv.Itoa(cfg.Port))
		}
		buf.WriteByte(')')
	}

	// /dbname
	buf.WriteByte('/')
	buf.WriteString(cfg.Name)

	return buf.String()
}

type MigrationConfig struct {
	Path string
}

type AuthConfig struct {
	JwtSecret      string
	JwtExpiryInSec int
	GoogleClientId string
}

type HttpConfig struct {
	ConnectTimeoutSec int
	RequestTimeoutSec int
	UserAgent         string
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
