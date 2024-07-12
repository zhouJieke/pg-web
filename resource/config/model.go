package config

import "time"

type App struct {
	Port       string `yaml:"port"`
	Version    string `yaml:"version"`
	Env        string `yaml:"env"`
	RunLogType string `yaml:"runLogType"`
}

type Log struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"`
	RootDir    string `mapstructure:"root_dir" json:"root_dir" yaml:"root_dir"`
	Filename   string `mapstructure:"filename" json:"filename" yaml:"filename"`
	Format     string `mapstructure:"format" json:"format" yaml:"format"`
	ShowLine   bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"` // MB
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`    // day
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}

type MysqlBase struct {
	Driver   string `mapstructure:"driver" json:"driver" yaml:"driver"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	UserName string `mapstructure:"userName" json:"userName" yaml:"userName"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
}

// RedisBase Redis配置项
type RedisBase struct {
	Address      string        `mapstructure:"address" json:"address" yaml:"address"`
	Password     string        `mapstructure:"password" json:"password" yaml:"password"`
	DB           int           `mapstructure:"db" json:"db" yaml:"db"`
	DialTimeout  time.Duration `mapstructure:"dialTimeout" json:"dialTimeout" yaml:"dialTimeout"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout" json:"readTimeout" yaml:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout" json:"writeTimeout" yaml:"writeTimeout"`
	PoolSize     int           `mapstructure:"poolSize" json:"poolSize" yaml:"poolSize"`
	PoolTimeout  time.Duration `mapstructure:"poolTimeout" json:"poolTimeout" yaml:"poolTimeout"`
}
