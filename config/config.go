package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

/*
Pada file ini hanya untuk mengambile file config untuk server dari yaml atau toml ke golang.

Kali ini kita menggunakan viper untuk mempermudah proses tersebut dan disini kita membuat beberapa struct untuk
memberikan kerapihan dalam struktur.

Terdapat banyak cara untuk membuat config database database ini. Salah satunya selain menggunakan toml atau yaml
adalah menggunakan .env. (Terserah Dev saja untuk ini)
*/
var (
	Viper *viper.Viper
)

type yamlConfig struct {
	DBConfig     DBConfig     `yaml:"database"`
	ServerConfig ServerConfig `yaml:"server"`
}

type DBConfig struct {
	Server        string `json:"server"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Port          string `json:"port"`
	Schema        string `json:"schema"`
	MaxConnection int    `json:"connection_max"`
}

type ServerConfig struct {
	Port  string `json:"port"`
	Debug string `json:"debug"`
}

var cfg yamlConfig

func Initialize(path string) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &cfg)
}

func Get(configName string) interface{} {
	if configName == "database" {
		return cfg.DBConfig
	} else if configName == "server" {
		return cfg.ServerConfig
	}
	return nil
}

func init() {
	readConfig("config/config")
}

func readConfig(filename string) {
	Viper = viper.New()
	Viper.AddConfigPath(".")
	Viper.SetConfigName(filename)
	err := Viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error when reading config")
	}
	//load from env variables
	replacer := strings.NewReplacer(".", "_")
	Viper.SetEnvKeyReplacer(replacer)
	Viper.AutomaticEnv()
}

func GetYamlValues() *yamlConfig {
	Db := &DBConfig{
		Server:        Viper.GetString("DATABASE.server"),
		Username:      Viper.GetString("DATABASE.username"),
		Password:      Viper.GetString("DATABASE.password"),
		Port:          Viper.GetString("DATABASE.port"),
		Schema:        Viper.GetString("DATABASE.schema"),
		MaxConnection: Viper.GetInt("DATABASE.connection_max"),
	}
	server := &ServerConfig{
		Port:  Viper.GetString("SERVICE.port"),
		Debug: Viper.GetString("SERVICE.debug"),
	}
	yml := &yamlConfig{*Db, *server}
	return yml
}
