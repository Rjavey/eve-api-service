package controller

import (
	"eve-api-service/log"
	"net/http"
)
import config "eve-api-service/conf"

func init() {

	http.HandleFunc("/hw", Helloworld)
	http.HandleFunc("/api/order/price", config.Search)

	var applicationConfig = config.ApplicationConfig

	log.Info.Printf("service is up on %v \n", applicationConfig.ServiceProt)
	err := http.ListenAndServe(":"+applicationConfig.ServiceProt, nil)
	if err != nil {
		log.Error.Printf("service handle err at ListenAndServe: ", err)
	}
}
