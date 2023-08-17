Translations: [English](README.md) | [简体中文](README_zh.md)

## 关于 go-comtrade

> go-comtrade是一个go语言库，用来解析COMTRADE录波文件.

## 开始

```shell
go get github.com/yonwoo9/go-comtrade
```

## 使用

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


## 开源协议

go-comtrade是基于[MIT license](./LICENSE)协议的开源软件.

## 致谢
以下这些库对go-comtrade开发设计有较大影响.

- [smslit/comtrade](https://github.com/smslit/comtrade).
- [dparrini/python-comtrade](https://github.com/dparrini/python-comtrade).
- [miguelmoreto/pycomtrade](https://github.com/miguelmoreto/pycomtrade).
- [ValleyZw/comgo](https://github.com/ValleyZw/comgo).

