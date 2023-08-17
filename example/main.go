package main

import (
	"flag"
	"fmt"
	"github.com/yonwoo9/go-comtrade"
)

var (
	configFile = flag.String("config", "testdata/test1.cfg", "config file path")
	dataFile   = flag.String("data", "testdata/test1.dat", "data file path")
)

func main() {
	flag.Parse()

	c, err := comtrade.ParseComtrade(*configFile, *dataFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(c.Conf)
}
