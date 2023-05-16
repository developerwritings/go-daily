package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	// Config
	viper.SetConfigName("default") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/") // config file path
	viper.AutomaticEnv()             // read value ENV variable

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	// Set default value
	viper.SetDefault("app.linetoken", "DefaultLineTokenValue")

	producerbroker := viper.GetString("app.producerbroker")
	consumerbroker := viper.GetString("app.consumerbroker")
	linetoken := viper.GetString("app.linetoken")

	fmt.Println("app.producerbroker :", producerbroker)
	fmt.Println("app.consumerbroker :", consumerbroker)
	fmt.Println("app.linetoken :", linetoken)
}
