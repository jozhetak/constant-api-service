package config

import (
	"encoding/json"
	"log"
	"os"
)

var config *Config

func init() {
	file, err := os.Open("conf/conf.json")
	if err != nil {
		log.Println("error:", err)
		panic(err)
	}
	decoder := json.NewDecoder(file)
	v := Config{}
	err = decoder.Decode(&v)
	if err != nil {
		log.Println("error:", err)
		panic(err)
	}
	config = &v
}

func GetConfig() *Config {
	return config
}

type Config struct {
	Port           int    `json:"port"`
	Db             string `json:"db"`
	TokenSecretKey string `json:"token_secret_key"`
	SendgridAPIKey string `json:"sendgrid_api_key"`
}
