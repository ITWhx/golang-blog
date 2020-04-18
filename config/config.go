/**
*
*@author 吴昊轩
*@create 2020-04-1517:31
 */
package config

import (
	"encoding/json"
	"os"
	"time"
)

var G_cfg *Config

type Config struct {
	ServerInfo serverInfo
	RedisInfo  redisInfo
}

type serverInfo struct {
	Host string
}

type redisInfo struct {
	Host        string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
	Db          int
}

func InitConfig() error {
	file, err := os.Open("config/config.json")
	defer file.Close()
	if err != nil {
		return err
	}
	config := Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}
	G_cfg = &config
	return nil
}
