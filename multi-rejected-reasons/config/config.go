package config

import (
	"fmt"
	"strings"

	"git.chotot.org/go-common/kit/logger"
	"github.com/spf13/viper"
)

type EchoServerConfig struct {
	Port  string `mapstructure:"port"`
	Debug bool   `mapstructure:"debug"`
}

type GrpcServerConfig struct {
	Network string `mapstructure:"network"`
	Address string `mapstructure:"address"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
}

// Schema schema
type Schema struct {
	Jwt        JWTConfig        `mapstructure:"jwt"`
	GrpcServer GrpcServerConfig `mapstructure:"grpc_server"`
	EchoServer EchoServerConfig `mapstructure:"echo_server"`
}

// ConfigMap configmap
var ConfigMap Schema
var log = logger.GetLogger("content-moderation-tools")

func init() {
	// Initialize viper default instance with API base config.
	Config := viper.New()
	Config.SetConfigName("config")       // Name of config file (without extension).
	Config.AddConfigPath("/etc/config/") // Optionally look for config in the working directory.
	Config.AddConfigPath(".")            // Look for config needed for tests.
	Config.AddConfigPath("./config")
	Config.AddConfigPath("./multi-rejected-reasons/config")
	// Config.AddConfigPath("../config")
	Config.AddConfigPath("../config/")

	Config.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))

	err := Config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	Config.AutomaticEnv()

	err = Config.Unmarshal(&ConfigMap)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error Unmarshal config file: %s", err))
	}
	log.Debugf("Current Config: %+v\n", ConfigMap)
}
