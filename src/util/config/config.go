package config

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBSource   string `mapstructure:"DB_SOURCE"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPass     string `mapstructure:"DB_PASSWORD"`
	DBUser     string `mapstructure:"DB_USER"`
	Collection string `mapstructure:"COLLECTION"`

	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	VaultHost     string `mapstructure:"VAULT_HOST"`
	VaultScheme   string `mapstructure:"VAULT_SCHEME"`
	VaultPort     string `mapstructure:"VAULT_PORT"`
	VaulToken     string `mapstructure:"VAULT_SECURITY_TOKEN"`
	ConsulPort    string `mapstructure:"CONSUL_PORT"`
	ConsulHost    string `mapstructure:"CONSUL_HOST"`
}

// GetConfig reads configuration from file or environment variables.
func GetConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
