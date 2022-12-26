package util

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`

	MessageMinSize int `mapstructure:"MESSAGE_MIN_SIZE"`
	MessageMaxSize int `mapstructure:"MESSAGE_MAX_SIZE"`

	ReceiverAddress    string `mapstructure:"RECEIVER_ADDRESS"`
	BrokerAddress      string `mapstructure:"BROKER_ADDRESS"`
	DestinationAddress string `mapstructure:"DESTINATION_SERVICE_ADDRESS"`

	BrokerLogDestination string `mapstructure:"BROKER_LOG_DESTINATION"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	// using $(path)/app.env as the config file
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// overwite config with environment variables if they exist
	viper.AutomaticEnv()

	// read config file from disk
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// unmarshal config file into Config struct
	err = viper.Unmarshal(&config)
	return
}
