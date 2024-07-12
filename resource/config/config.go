package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	App         App       `mapstructure:"app" json:"app" yaml:"app"`
	Log         Log       `mapstructure:"log" json:"log" yaml:"log"`
	MysqlConfig MysqlBase `mapstructure:"mysqlConfig" json:"mysqlConfig" yaml:"mysqlConfig"`
	RedisConfig RedisBase `mapstructure:"redis" json:"redis" yaml:"redis"`
}

func (config *Config) InitializeConfig() *Config {
	//获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	vip := viper.New()
	vip.AddConfigPath(path + "/resource") //设置读取的文件路径
	vip.SetConfigName("config")           //设置读取的文件名
	vip.SetConfigType("yaml")             //设置文件的类型
	//尝试进行配置读取
	if err := vip.ReadInConfig(); err != nil {
		panic(err)
	}
	// 监听配置文件
	vip.WatchConfig()

	vip.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := vip.Unmarshal(&config); err != nil {
			fmt.Println(err)
		}
	})
	err = vip.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return config
}
