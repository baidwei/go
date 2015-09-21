package main

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://api.tianditu.com/api/api-new/BuslineServlet.do?postStr={'startposition':'111.69928,40.78133','endposition':'111.70329,40.77628',linetype:'1'}")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	/*
		var v interface{}
		errs := json.Unmarshal(body, &v)
		if errs != nil {
			fmt.Println(errs)
		}
		fmt.Println(v)
	*/
}
