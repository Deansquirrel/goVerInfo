package common

import (
	"fmt"
	"github.com/Deansquirrel/go-tool"
	"github.com/Deansquirrel/goVerInfo/global"
)

func PrintAndLog(msg string) {
	if !global.SysConfig.Total.IsDebug {
		err := go_tool.Log(msg)
		if err != nil {
			fmt.Println("write log error:" + err.Error())
			fmt.Println("msg:" + msg)
		}
	} else {
		fmt.Println(msg)
	}

}
