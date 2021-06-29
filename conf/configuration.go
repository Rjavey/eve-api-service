package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	RedisHost string
	RedisProt string

	AccessToken  string
	RefreshToken string

	ServiceProt string
}

var ApplicationConfig Config

func init() {
	log.Println("start init config")
	bytes, err := ioutil.ReadFile("./config.json")
	ApplicationConfig = Config{RedisHost: "localhost", RedisProt: "6379", ServiceProt: "22022"}
	err = json.Unmarshal(bytes, &ApplicationConfig)
	if err != nil {
		log.Fatalf("err was %v", err)
	}
}
