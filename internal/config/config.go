package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"github.com/tonghia/go-challenge-transaction-app/pkg/server"
)

// Config holds all settings
//
//go:embed config.yaml
var defaultConfig []byte

type Config struct {
	Base  `mapstructure:",squash"`
	MySQL *MySQL `yaml:"mysql" mapstructure:"mysql"`
}

type Base struct {
	Server      ServerConfig `yaml:"server" mapstructure:"server"`
	Environment string       `yaml:"environment" mapstructure:"environment"`
}

// MySQL is settings of a MySQL server. It contains almost same fields as mysql.Config,
// but with some different field names and tags.
type MySQL struct {
	Username string `yaml:"username" mapstructure:"username"`
	Password string `yaml:"password" mapstructure:"password"`
	Protocol string `yaml:"protocol" mapstructure:"protocol"`
	Address  string `yaml:"address" mapstructure:"address"`
	Port     int    `yaml:"port" mapstructure:"port"`
	Database string `yaml:"database" mapstructure:"database"`
}

// FormatDSN returns MySQL DSN from settings.
func (m *MySQL) FormatDSN() string {
	um := &mysql.Config{
		User:   m.Username,
		Passwd: m.Password,
		Net:    m.Protocol,
		Addr:   m.Address,
		DBName: m.Database,
	}
	return um.FormatDSN()
}

// ServerConfig hold http/grpc server config
type ServerConfig struct {
	GRPC server.Listen `json:"grpc" mapstructure:"grpc" yaml:"grpc"`
	HTTP server.Listen `json:"http" mapstructure:"http" yaml:"http"`
}

func Load() (*Config, error) {
	cfg := new(Config)

	viper.SetConfigType("yaml")

	if err := viper.ReadConfig(bytes.NewBuffer(defaultConfig)); err != nil {
		return nil, fmt.Errorf("Failed to read viper config: %v", err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal config: %v", err)
	}

	return cfg, nil
}
