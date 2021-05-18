package main

import (
	"github.com/jinzhu/configor"
	"log"
)

//Config contains global app's configuration
var Config AppConfig

type User struct {
	Name   string
	Groups []string
	Key    string
}

type InfoSource struct {
	ID   string
	Name string
	Exec string
}

type CommandTarget struct {
	ID      string
	Name    string
	Details string
	Exec    []string
	Groups  []string
	Danger  bool
}

// AppConfig contains app's configuration
type AppConfig struct {
	Commands []CommandTarget
	Info     []InfoSource
	Users    []User

	Key    string
	Port   string
	Cors   string
	Server string
	Seed   string
}

//LoadFromFile method loads and parses config file
func (c *AppConfig) LoadFromFile(url string) {
	err := configor.Load(&Config, url)
	if err != nil {
		log.Fatalf("Can't load the config file: %s", err)
	}
}
