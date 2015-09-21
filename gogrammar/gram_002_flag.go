package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	var confFile string
	flag.StringVar(&confFile, "conf", "app.conf", "")
	flag.Parse()

	//fmt.Println(confFile)
	//flag.PrintDefaults()
	b, e := ioutil.ReadFile(confFile)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(b))
}
