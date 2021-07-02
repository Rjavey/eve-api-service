package conf

import (
	"encoding/json"
	"eve-api-service/entry"
	"eve-api-service/log"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const ORDER_PRE = "TypeOrder:"
const JITA = "10000002:"
const TYPE_HOT = "TypeHot:"

const BASE_URL = "https://esi.evetech.net/latest"
const ORDER = "/markets/%s/orders/"

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
		result := Get(ORDER_PRE + JITA + param.TypeId)
		if len(result) != 0 {
			err = json.Unmarshal([]byte(result), &orders)
		}
		break
	case len(param.TypeName) != 0:
		// todo
		keys := Scan(ORDER_PRE + JITA + param.TypeName)
		for _, i := range keys {
			result := Get(ORDER_PRE + JITA + i)
			if len(result) != 0 {
				err = json.Unmarshal([]byte(result), &orders)
			}
		}
		break
	default:
		fmt.Fprintf(w, "can't find any search param")
	}

	var order entry.Order

	// 后续处理
	Incr(TYPE_HOT + order.TypeId)

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

func FindOrder(id string, buy string, regin string) {
	params := url.Values{}
	str := fmt.Sprintf(BASE_URL+ORDER, regin)
	Url, _ := url.Parse(str)
	params.Set("order_type ", buy)
	params.Set("type_id", id)
	urlPath := Url.String()
	//fmt.Println(urlPath)
	resp, _ := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
