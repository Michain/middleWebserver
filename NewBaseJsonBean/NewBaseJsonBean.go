package NewBaseJsonBean

import (
	"io"
	"webserver/serialization"
)

//Server Config json for struct type
//定义服务器配置文件结构（json）
type ServerConfigBean struct {
	HostIp  string `json:"HostIp"`
	APIPort string `json:"APIPort"`
}

//Input json interface for struct type
//定义输入接口结构（json）
type login struct {
	UserName string `json:"userName"`
	PassWord string `json:"password"`
}

//Output json interface for struct type
//定义输出接口结构（json）
type BaseJsonBean struct {
	Code    int         `json:"code"`
	Header  string      `json:"header"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type Input struct {
	Identity string `json:"identity"`
}

type OutputDate struct {
	Identity string `json:"identity"`
	UserName string `json:"userName"`
	Money    string `json:"money"`
	Txtype   string `json:"txtype"`
	Date     string `json:"date"`
}

func NewInputJsonBean() *Input {
	return &Input{}
}

func NewOutputJsonBean() *BaseJsonBean {
	return &BaseJsonBean{}
}

func NewServerConfigBean() *ServerConfigBean {
	return &ServerConfigBean{}
}
func (a *OutputDate) Serialize(w io.Writer) error {
	serialization.WriteVarString(w, a.Identity)
	serialization.WriteVarString(w, a.UserName)
	serialization.WriteVarString(w, a.Money)
	serialization.WriteVarString(w, a.Txtype)
	serialization.WriteVarString(w, a.Date)
	return nil
}

func (a *OutputDate) Deserialize(w io.Reader) error {
	a.Identity, _ = serialization.ReadVarString(w)
	a.UserName, _ = serialization.ReadVarString(w)
	a.Money, _ = serialization.ReadVarString(w)
	a.Txtype, _ = serialization.ReadVarString(w)
	a.Date, _ = serialization.ReadVarString(w)

	return nil
}
