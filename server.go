package main

import (
	"fmt"
	"github.com/Deansquirrel/goVerInfo/common"
	"github.com/Deansquirrel/goVerInfo/global"
	"github.com/Deansquirrel/goVerInfo/webService"
	"github.com/kardianos/service"
	"os"
)

type program struct {
}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func (p *program) run() {
	ws := webService.WebService{
		Port: global.SysConfig.Total.Port,
	}
	ws.Start()
}

func main() {
	err := global.RefreshConfig()
	if err != nil {
		common.PrintAndLog("加载配置时遇到错误:" + err.Error())
		return
	}

	svcConfig := &service.Config{
		Name:        "Name",        //服务显示名称
		DisplayName: "DisplayName", //服务名称
		Description: "Description", //服务描述
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		common.PrintAndLog("创建服务对象时遇到错误:" + err.Error())
		return
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			installService(s)
			return
		case "uninstall":
			unInstallService(s)
			return
		default:
			fmt.Println("未识别的参数名称\n安装服务:install\n卸载服务:uninstall")
			return
		}
	}

	err = s.Run()
	if err != nil {
		common.PrintAndLog("运行时遇到错误:" + err.Error())
	}
}

func installService(s service.Service) {
	err := s.Install()
	msg := ""
	if err != nil {
		msg = "安装服务时遇到错误:" + err.Error()
	} else {
		msg = "服务安装成功"
	}
	fmt.Println(msg)
}

func unInstallService(s service.Service) {
	err := s.Uninstall()
	msg := ""
	if err != nil {
		msg = "卸载服务时遇到错误:" + err.Error()
	} else {
		msg = "服务卸载成功"
	}
	fmt.Println(msg)
}
