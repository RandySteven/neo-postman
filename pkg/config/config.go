package config

import "time"

type (
	Config struct {
		Server struct {
			Host    string `yaml:"host"`
			Port    string `yaml:"port"`
			Timeout struct {
				Server time.Duration `yaml:"server"`
				Read   time.Duration `yaml:"read"`
				Write  time.Duration `yaml:"write"`
				Idle   time.Duration `yaml:"idle"`
			} `yaml:"timeout"`
		} `yaml:"server"`

		Postgres struct {
			Host   string `yaml:"host"`
			Port   string `yaml:"port"`
			DbName string `yaml:"dbname"`
			DbUser string `yaml:"dbuser"`
			DbPass string `yaml:"dbpass"`
		} `yaml:"postgres"`

		Redis struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Password string `yaml:"password"`
			DB       int    `yaml:"db"`
		} `yaml:"redis"`
	}
)
