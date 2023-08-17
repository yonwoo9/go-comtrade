
##### Translate to: [简体中文](README_zh.md)
## About go-comtrade
  
> The go-comtrade is a lib used to parse COMTRADE file.

## Getting Started
 
```shell
go get github.com/yonwoo9/go-comtrade
```

## Usage
```go
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
```


## License

The go-comtrade is open-sourced software licensed under the [MIT license](./LICENSE).
  
## Acknowledgments
  
The following project had particular influence on go-comtrade design.
  
- [smslit/comtrade](https://github.com/smslit/comtrade).
- [dparrini/python-comtrade](https://github.com/dparrini/python-comtrade).
- [miguelmoreto/pycomtrade](https://github.com/miguelmoreto/pycomtrade).
- [ValleyZw/comgo](https://github.com/ValleyZw/comgo).
