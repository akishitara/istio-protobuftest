package config

import (
	"github.com/BurntSushi/toml"
)

// Config load config.toml
type SetupConfig struct {
	Server ServerConfig
	Client ClientConfig
	Redis  RedisConfig
}

// ServerConfig load config toml [server]
type ServerConfig struct {
	Address string
	Port    int
}

// ClientConfig load config toml [client]
type ClientConfig struct {
	Address string
	Port    int
}

// RedisConfig load config toml [client]
type RedisConfig struct {
	Address string
	Port    int
}

// ConfigLoad load toml
func ConfigLoad(filePath string) SetupConfig {
	var conf SetupConfig
	_, err := toml.DecodeFile(filePath, &conf)
	if err != nil {
		panic(err)
	}
	return conf
}
