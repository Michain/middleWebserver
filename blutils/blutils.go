package blutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	. "webserver/NewBaseJsonBean"
	. "webserver/commonutils"
)

func LoginTask(w http.ResponseWriter, req *http.Request) {
	fmt.Println("loginTask is running...")
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(req.Body)
	DebugHttpHeader("loginTask", body, req)

	out := make(map[string]interface{})
	out["userName"] = ""
	out["password"] = ""

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
	time.Sleep(time.Second * 2)
	data := OutputDetailSet(searchReq.Identity)
	result := ReplySet(data)

	//向客户端返回JSON数据
	bytes, _ := json.Marshal(result)
	fmt.Fprint(w, string(bytes))
}

func OutputDetailSet(identity string) []byte {
	b_buf := new(bytes.Buffer)
	out := OutputDate{}

	out.Identity = identity
	if       identity == "111111111111111111"{
		out.UserName = "zhangsan"
	}else if identity == "222222222222222222"{
		out.UserName = "lisi"
	}else if identity == "333333333333333333"{
		out.UserName = "zhaoliu"
	}else{
		out.UserName = "huahua"
	}
	out.Money = "30000"
	out.Txtype = "typex"
	out.Date = "2017/01/27"

	out.Serialize(b_buf)

	return b_buf.Bytes()
}

func ReplySet(data []byte) BaseJsonBean {
	result := BaseJsonBean{}
	result.Code = 100
	result.Header = "Content-Type, application/json;charset=UTF-8"
	result.Data = string(data)
	result.Message = "Sucessful"

	return result
}
