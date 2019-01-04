package global

import (
	"errors"
	"github.com/BurntSushi/toml"
	"github.com/Deansquirrel/go-tool"
	"github.com/Deansquirrel/goVerInfo/object"
)

var SysConfig object.SysConfig

func RefreshConfig() error {
	err := refreshSysConfig()
	if err != nil {
		return errors.New("刷新SysConfig:" + err.Error())
	}
	return nil
}

func refreshSysConfig() error {
	path, err := go_tool.GetCurrPath()
	if err != nil {
		return err
	}

	var sysConfig object.SysConfig
	_, err = toml.DecodeFile(path+"\\"+"config.toml", &sysConfig)
	if err != nil {
		return nil
	}
	SysConfig = sysConfig
	return nil
}
