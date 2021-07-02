package conf

import (
	"eve-api-service/log"
	"time"
)

var typeIdCh = make(chan string, 100)

/**
启动定时任务
*/
func InitJob() {

	ticker := time.NewTicker(time.Second * 10)
	go func() {
		for _ = range ticker.C {
			log.Info.Printf("ticked job start")
			var typeId string
			typeIdCh <- typeId

			FindOrder("804", "buy", "10000002")

		}
	}()
}
