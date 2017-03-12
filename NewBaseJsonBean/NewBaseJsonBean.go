package NewBaseJsonBean

//Server Config json for struct type
//定义服务器配置文件结构（json）
type ServerConfigBean struct {
	HostIp  string `json:"HostIp"`
	APIPort string `json:"APIPort"`
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
	SerialNumber string `json:"SerialNumber"`
	Trstype string `json:"Trstype"`
	TimeStamp string `json:"TimeStamp"`
}

func NewInputJsonBean() *Input {
	return &Input{}
}

func NewServerConfigBean() *ServerConfigBean {
	return &ServerConfigBean{}
}

type OutputDate struct {
	Response string `json:"Response"`
}