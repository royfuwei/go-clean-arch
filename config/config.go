package config

import (
	"flag"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Configs struct {
	Port          string
	RetryPeriod   int
	MaxRetryCount int
	Goroutine     int
	MongoAddr     string
	MgoDBName     string
}

var (
	Cfgs *Configs
)

func initVariable() {
	flag.Set("alsologtostderr", "true")
	flag.CommandLine.Parse([]string{})
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.String("port", "5001", "serve port")
	pflag.Int("retry_period", 3, "second of retry period")
	pflag.Int("max_retry_count", 5, "max retry count")
	pflag.Int("goroutine", 1, "goroutine number")
	pflag.String("mongo_addr", "127.0.0.1:27017", "mongodb address")
	pflag.String("mongo_db_name", "core", "mongodb name")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

// NewConfig new config
func NewConfig() *Configs {
	initVariable()
	Cfgs = &Configs{
		Port:          ":" + viper.GetString("port"),
		RetryPeriod:   viper.GetInt("retry_period"),
		MaxRetryCount: viper.GetInt("max_retry_count"),
		Goroutine:     viper.GetInt("goroutine"),
		MongoAddr:     viper.GetString("mongo_addr"),
		MgoDBName:     viper.GetString("mongo_db_name"),
	}
	// initConfigsJson(Cfgs.ConfigsInitPath)
	return Cfgs
}
