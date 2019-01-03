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
	//err := loadConfig()
	//if err != nil {
	//	common.PrintAndLog(err.Error())
	//	return
	//}
	//common.PrintAndLog("service run")

	err := loadConfig()
	if err != nil {
		common.PrintAndLog("加载配置文件时遇到错误:" + err.Error())
		return
	}

	ws := webService.WebService{
		Port: global.SysConfig.Total.Port,
	}
	ws.Start()
}

func main() {
	//==============================================================================
	//_ = loadConfig()
	//
	//ws := webService.WebService{
	//	Port:global.SysConfig.Total.Port,
	//}
	//go func(){
	//	err := ws.Start()
	//	if err != nil {
	//		common.PrintAndLog(err.Error())
	//	}
	//}()
	//==============================================================================

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			installService()
			return
		case "uninstall":
			unInstallService()
			return
		default:
			fmt.Println("未识别的参数名称\n安装服务:install\n卸载服务:uninstall")
			return
		}
	} else {
		p := program{}
		p.run()
	}
}

func loadConfig() error {
	err := global.RefreshConfig()
	if err != nil {
		common.PrintAndLog(err.Error())
		return err
	}
	return nil
}

func getService() (service.Service, error) {
	svcConfig := &service.Config{
		Name:        "Name",
		DisplayName: "DisplayName",
		Description: "Description",
	}

	prg := &program{}
	return service.New(prg, svcConfig)
}

func installService() {
	s, err := getService()
	if err != nil {
		common.PrintAndLog("Install error:" + err.Error())
		return
	}
	err = s.Install()
	msg := ""
	if err != nil {
		msg = err.Error()
	} else {
		msg = "服务安装成功"
	}
	fmt.Println(msg)
	common.PrintAndLog(msg)
}

func unInstallService() {
	s, err := getService()
	if err != nil {
		common.PrintAndLog("UnInstall error:" + err.Error())
		return
	}
	err = s.Uninstall()
	msg := ""
	if err != nil {
		msg = err.Error()
	} else {
		msg = "服务卸载成功"
	}
	fmt.Println(msg)
	common.PrintAndLog(msg)
}
