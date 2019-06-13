package cmd

import (
	"os"

	"github.com/spf13/viper"
)

var (
	WorkingPath string
	LogFile     string
	RpcAddr     string
	SavingData  string
)

func SetupConfig() {
	homeDir := os.Getenv("HOME")
	WorkingPath = homeDir + "/.gojuon"
	if _, err := os.Stat(WorkingPath); os.IsNotExist(err) {
		if err := os.Mkdir(WorkingPath, 0744); err != nil {
			panic(err)
		}
	}

	needWrite := false
	configFile := WorkingPath + "/gojuon.yaml"
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		needWrite = true
		if _, err := os.Create(configFile); err != nil {
			panic(err)
		}
	}
	viper.AddConfigPath(WorkingPath)
	viper.SetConfigFile("yaml")
	viper.SetConfigName("gojuon")
	viper.SetDefault("global.rpc_addr", ":20443")
	viper.SetDefault("global.log_file", "/gojuon.log")
	viper.SetDefault("global.data", WorkingPath+"/data")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	LogFile = WorkingPath + viper.GetString("global.log_file")
	RpcAddr = viper.GetString("global.rpc_addr")
	SavingData = viper.GetString("global.data")

	if needWrite {
		if err := viper.WriteConfigAs(configFile); err != nil {
			panic(err)
		}
	}
}
