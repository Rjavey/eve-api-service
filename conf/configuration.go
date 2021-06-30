package conf

import (
	"encoding/json"
	"eve-api-service/log"
	"io/ioutil"
)

type Config struct {
	RedisHost string
	RedisProt string
	RedisAuth string

	AccessToken  string
	RefreshToken string

	ServiceProt string
}

var ApplicationConfig Config

func init() {
	log.Info.Println("start init config")
	bytes, err := ioutil.ReadFile("./config.json")
	ApplicationConfig = Config{RedisHost: "localhost", RedisProt: "6379", ServiceProt: "22022"}
	err = json.Unmarshal(bytes, &ApplicationConfig)
	if err != nil {
		log.Error.Printf("err was %v", err)
	}

	// init redis pool
	connRedis(ApplicationConfig.RedisHost+":"+ApplicationConfig.RedisProt, ApplicationConfig.RedisAuth)
}
