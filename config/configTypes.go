package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type DBConfig struct {
	DataDir           string `mapstructure:"dataDir"`
	StorageFile       string `mapstructure:"storageFile"`
	IgnoreVersionTest bool   `mapstructure:"ignoreVersionTest"`
	EnableDebug       bool   `mapstructure:"enableDebug"`
}

type Config struct {
	// The name of the service
	//dataDir string `mapstructure:"datadir" default:"./data" yaml:"datadir" json:"datadir"`
	//DataDir  string   `mapstructure:"datadir"`
	DBConfig DBConfig `mapstructure:"dbConfig"`
}

// CONFIG is the configuration for the service
var CONFIG Config

// LoadConfig loads the configuration
func LoadConfig() error {

	// Load the configuration using viper
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	//viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	//viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath("./") // optionally look for config in the working directory
	path, _ := os.Getwd()
	viper.AddConfigPath(path)   // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}

	// Unmarshal the configuration into the CONFIG variable
	cfg := Config{}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return fmt.Errorf("unable to decode into struct, %v", err)
	}
	//fmt.Println(cfg.DBConfig.DataDir + " 888")

	err = viper.Unmarshal(&CONFIG)
	if err != nil {
		return fmt.Errorf("unable to decode into struct, %v", err)
	}
	fmt.Println(os.Getwd())
	data, er := json.Marshal(CONFIG)
	fmt.Println(er)
	fmt.Println(string(data))
	fmt.Println(viper.Get(("datadir")))
	return nil
}
