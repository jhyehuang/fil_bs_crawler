package config

import (
	"github.com/fil_bs_crawler/internal/flogging"
	"github.com/spf13/viper"
)

var SingleViper *viper.Viper
var ChainViper *viper.Viper

func init() {
	//ret, _ := os.Getwd()
	//fmt.Printf("-------------%s\n", ret)
	SingleViper = viper.New()
	SingleViper.SetConfigFile("/Users/huangzhijie/code/go-project/sxxl_code/fil_bs_crawler/config/config.toml")
	err := SingleViper.ReadInConfig()
	if err != nil {
		flogging.Log.Fatalf("[init] read config.toml error: %s\n", err)
	}

	// init chain config
	ChainViper = viper.New()
	ChainViper.SetConfigFile("/Users/huangzhijie/code/go-project/sxxl_code/fil_bs_crawler/config/chain.toml")
	err = ChainViper.ReadInConfig()
	if err != nil {
		flogging.Log.Error("[init] init chain config error: %s", err)
		panic(err)
	}
}

func GetHttpListenAddress() string {
	return SingleViper.GetString("http.address")
}

type LoggerConfig struct {
	FileName   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

func GetLogger() *LoggerConfig {
	lc := &LoggerConfig{
		FileName:   SingleViper.GetString("logger.FileName"),
		MaxSize:    SingleViper.GetInt("logger.MaxSize"),
		MaxBackups: SingleViper.GetInt("logger.MaxBackups"),
		MaxAge:     SingleViper.GetInt("logger.MaxAge"),
		Compress:   SingleViper.GetBool("logger.Compress"),
	}
	return lc
}
