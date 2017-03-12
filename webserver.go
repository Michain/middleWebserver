package main

import (
	"fmt"
	"net/http"
	. "middleWebserver/conf"
	. "middleWebserver/router"
)

func main() {
	//Based on the business request, interface route to the function
	//到router的程序代码添加业务名和对应的处理函数设定
	Httpreqrouter()

	//Read the server config file,get host ip and port
	//读取服务器配置文件设定要监听的主机地址和端口号
	httpserver, err := GetServerConfig()
	if err != nil {
		fmt.Printf("[httpserver config error] %s\n", err.Error())
	}

	//Listen Http server's hostip and port for request
	//服务器监听地址和端口号
	fmt.Printf("http server listening at %s end\n", httpserver)
	err = http.ListenAndServe(":5000", nil)
	//catch http server startup exception
	//服务器监听异常捕获处理
	if err != nil {
		fmt.Printf("ListenAndServe error: %s end.\n", err.Error())
	}

}
