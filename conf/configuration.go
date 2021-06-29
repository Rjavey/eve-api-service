package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	fmt.Println("start init config")
	bytes, err := ioutil.ReadFile("./config.json")
	ApplicationConfig = Config{RedisHost: "localhost", RedisProt: "6379", ServiceProt: "9988"}
	err = json.Unmarshal(bytes, &ApplicationConfig)
	if err != nil {
		fmt.Printf("err was %v", err)
	}
	fmt.Printf("redis uri is %v:%v \n", ApplicationConfig.RedisHost, ApplicationConfig.RedisProt)
	fmt.Printf("service prot is %v \n", ApplicationConfig.ServiceProt)
}
