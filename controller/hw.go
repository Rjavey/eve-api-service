package controller

import (
	"fmt"
	"net/http"
)

func Helloworld(w http.ResponseWriter, r *http.Request) {
	// 往w里写入内容，就会在浏览器里输出
	fmt.Fprintf(w, "service is up!")
}
