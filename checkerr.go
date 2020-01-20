package main

import (
	"fmt"
	"runtime"
)

func CheckErr(err error) bool {
	var file string
	var line int
	var ok bool
	if err != nil {
		_, file, line, ok = runtime.Caller(1)
		if ok {
			emsg := fmt.Sprintf("file:%s, line:%d, error:%s", file, line, err.Error())
			fmt.Println(emsg)
		} else {
			fmt.Println(err)
		}
		return true
	}
	return false
}
