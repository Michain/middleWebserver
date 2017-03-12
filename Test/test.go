package main

import (
"fmt"
"strings"
"net/http"
"io/ioutil"
)

func main() {
	fmt.Println(PostComplete("123456","1","2017/03/11:300000"))
}

func PostComplete(SerialNumber string,Trstype string,TimeStamp string) string{
	url := "http://35.185.139.212:5000/Complete"


	//payload := strings.NewReader("{\r\n  \"SerialNumber\": \"123456\",\r\n  \"Trstype\": \"1\",\r\n  \"TimeStamp\": \"2017/03/11:300000\"\r\n}")
	payload := strings.NewReader("{\r\n  \"SerialNumber\": \"" + SerialNumber + "\",\r\n  \"Trstype\": \""+Trstype+"\",\r\n  \"TimeStamp\": \""+TimeStamp+"\"\r\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "bf098433-e4ba-7716-922b-4c21cd0a33eb")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return  string(body)
}