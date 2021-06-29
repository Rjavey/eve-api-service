package controller

import (
	"log"
	"net/http"
)
import config "eve-api-service/conf"

func init() {

	http.HandleFunc("/hw", Helloworld)
	var config = config.ApplicationConfig

	log.Printf("service is up on %v \n", config.ServiceProt)
	err := http.ListenAndServe(":"+config.ServiceProt, nil)
	if err != nil {
		log.Fatal("service handle err at ListenAndServe: ", err)
	}
}
