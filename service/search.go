package service

import (
	"encoding/json"
	redisUtil "eve-api-service/conf"
	"eve-api-service/entry"
	"eve-api-service/log"
	"fmt"
	"net/http"
)

const ORDER_PRE = "TypeOrder:"
const JITA = "100001:"

/**
查询
*/
func Search(w http.ResponseWriter, r *http.Request) {

	isGet(w, r)

	query := r.URL.Query()
	jsonStr, err := json.Marshal(query)
	if err != nil {
		log.Error.Printf("parse query param err: %v\n", err)
	}
	var param entry.TypeSearchParam
	err = json.Unmarshal(jsonStr, &param)
	if err != nil {
		log.Error.Printf("parse json err was %v", err)
	}

	var orders []entry.TypeSearchResult
	switch {
	case len(param.TypeId) != 0:
		result := redisUtil.Get(ORDER_PRE + JITA + param.TypeId)
		if len(result) != 0 {
			err = json.Unmarshal([]byte(result), &orders)
		}
		break
	case len(param.TypeName) != 0:
		// todo
		keys := redisUtil.Scan(ORDER_PRE + JITA + param.TypeName)
		for _, i := range keys {
			result := redisUtil.Get(ORDER_PRE + JITA + i)
			if len(result) != 0 {
				err = json.Unmarshal([]byte(result), &orders)
			}
		}
		break
	default:
		fmt.Fprintf(w, "can't find any search param")
	}
	fmt.Fprintf(w, "todo")

}

/**
更新缓存
*/
func UpdateCache(order entry.Order) {

}

/**
检查请求类型
*/
func isGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Error.Printf("invalid_http_method:%v", r)
	}
}
