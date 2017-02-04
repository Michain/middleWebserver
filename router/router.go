package router

import (
	"net/http"
	. "webserver/blutils"
)

func Httpreqrouter() {
	//the first parameter is http request interface,
	//the second parameter is th function that does with business logic
	//第一个参数为客户端发起http请求时的接口名，第二个参数是一个func，负责处理这个请求。
	http.HandleFunc("/login", LoginTask)
}
