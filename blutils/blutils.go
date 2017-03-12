package blutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	. "middleWebserver/NewBaseJsonBean"
	. "middleWebserver/commonutils"
)

func TransactionComplete(w http.ResponseWriter, req *http.Request) {
	fmt.Println("TransactionComplete is running...")
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)
	DebugHttpHeader("TransactionComplete", body, req)

	err := CheckHttpMethod(req, "POST")
	if err != nil {
		fmt.Printf("[loginTask]Wrong POST Method %s\n", err)
		return
	}

	//  searchReq := NewBaseJsonBean.NewInputJsonBean()
	searchReq := NewInputJsonBean()
	err = json.Unmarshal(body, searchReq)
	if err != nil {
		fmt.Printf("[loginTask]Wrong POST Data:["+string(body)+"]", err)
		return
	}
	//模拟延时
	time.Sleep(time.Second * 1)
	fmt.Println(searchReq)
	result := ReplySet()

	//向客户端返回JSON数据
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
}

func ReplySet() BaseJsonBean {
	result := BaseJsonBean{}
	result.Code = 000
	result.Header = "Content-Type, application/json;charset=UTF-8"
	result.Data = "OK"
	result.Message = "Sucessful"

	return result
}