package config

import (
	"fmt"
	"strings"

	"git.chotot.org/go-common/kit/logger"
	"github.com/spf13/viper"
)

// Schema schema
type Schema struct {
	App struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"app"`
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
