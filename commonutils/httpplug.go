package commonutils

import (
    "fmt"
    "net/http"
    "errors"
)

func DebugHttpHeader(methodName string, body []byte, req *http.Request) {
	method := req.Method
	userAgent := req.UserAgent()
	requsetUri := req.RequestURI
	contentType := req.Header.Get("Content-Type")
	fmt.Printf("[%s] Executed.Method=[%s],UserAgent=[%s],ContentType=[%s],RequestURI=[%s],Body=[%s]\n", methodName, method, userAgent, contentType, requsetUri, body)
}

func CheckHttpMethod(req *http.Request, targetMethod string) error {
	method := req.Method
	if method != targetMethod {
		msg := fmt.Sprintf("Wrong Http Method.Expect:%s,Actual:%s", targetMethod, method)
		return errors.New(msg)
	}
	return nil
}
