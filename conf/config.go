package cofig

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	. "middleWebserver/NewBaseJsonBean"
)

func GetServerConfig() (string, error) {
	var result string
	result = ""
	//Set the server config file path
	//指定服务器配置文件路径并读取服务器配置参数【jsson格式】
	//ServerConfigFile :="/opt/gopath/src/github.com/hyperledger/fabric/webserver/conf/config.json"
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	fmt.Printf("[Path=]%s", path)
	rs := []rune(path)
	rl := len(rs) - 10
	configpath := string(rs[0:rl])
	configfile := "/conf/config.json"
	ServerConfigFile := configpath + configfile
	fmt.Printf("[Config file ServerConfigFile]%s\n", ServerConfigFile)
	bytes, err := ioutil.ReadFile(ServerConfigFile)
	//catch server config file read exception
	//文件读取异常捕获处理
	if err != nil {
		msg := fmt.Sprintf("[Config file read error]%s", err.Error())
		return result, errors.New(msg)
	}

	//New ServerConfigBean struct type
	//创建ServerConfigBean结构体对象
	Server := NewServerConfigBean()

	//Unmarshl serverconfig file json to ServerConfigBean struct object
	//解析服务器配置文件到结构体对象
	err = json.Unmarshal(bytes, &Server)
	//catch json Unmarshal exception
	//文件内容解析异常捕获处理
	if err != nil {
		msg := fmt.Sprintf("[Config file Unmarshal error]%s", err.Error())
		return result, errors.New(msg)
	}

	//Edit the host ip and port
	//服务器ip地址和端口拼接
	result = Server.HostIp + ":" + Server.APIPort

	//normal return functiion parameter
	//正常返回
	return result, nil
}
