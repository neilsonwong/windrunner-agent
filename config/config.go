package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

// Config holds the go model for the config json file
type Config struct {
	Version         string `mapstructure:"version" json:"version"`
	ShareServer     string `mapstructure:"share_server" json:"share_server"`
	ShareServerAddr string `mapstructure:"share_server_addr" json:"share_server_addr"`
	ShareFolder     string `mapstructure:"share_folder" json:"share_folder"`
	ListingServer   string `mapstructure:"listing_server" json:"listing_server"`
	OsxMountPoint   string `mapstructure:"osx_mount" json:"osx_mount"`
	ServerPort      int    `mapstructure:"server_port" json:"server_port"`
}

// Load loads the config into viper and returns and config object
func Load() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	return Get()
}

// Print prints the currently loaded configuration parameters
func (config *Config) Print() {
	// print our config
	log.Println("config: share located at " + "//" + config.ShareServer + "/" + config.ShareFolder)
	log.Println("config: listing server at " + config.ListingServer)
	log.Println("config: osx mount point at " + config.OsxMountPoint)
	log.Println("config: agent hosted on port " + strconv.Itoa(config.ServerPort))
}

// Get the config that is being used
func Get() *Config {
	conf := &Config{}
	err := viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}
	return conf
}

// Update updates the config file and reloads configurations
func Update(updates *Config) {
	fmt.Println(updates)
	// backup current config
	viper.WriteConfigAs("./config.json." + strconv.FormatInt(time.Now().Unix(), 10))

	// overwrite current config file
	file, _ := json.MarshalIndent(updates, "", "    ")
	err := ioutil.WriteFile("config.json", file, 0755)

	if err != nil {
		log.Printf("error writing config file\n%s", err.Error())
	}

	// only reloads the config, wont restart the application
	Load()
}
