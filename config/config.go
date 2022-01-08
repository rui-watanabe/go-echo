package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	DBUser string
	DBPass string
	DBHost string
	DBName string
	DBPort int
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to readfile :%v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		DBUser: cfg.Section("db").Key("user").String(),
		DBPass: cfg.Section("db").Key("pass").String(),
		DBHost: cfg.Section("db").Key("host").String(),
		DBName: cfg.Section("db").Key("name").String(),
		DBPort: cfg.Section("db").Key("port").MustInt(),
	}
}
