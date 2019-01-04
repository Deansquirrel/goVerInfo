package webService

import (
	"fmt"
	"github.com/Deansquirrel/goVerInfo/common"
)

type IrisExternalLogger struct {
}

func (logger *IrisExternalLogger) Print(a ...interface{}) {
	common.PrintAndLog(fmt.Sprint(a))
}

func (logger *IrisExternalLogger) Println(a ...interface{}) {
	common.PrintAndLog(fmt.Sprint(a))
}

func (logger *IrisExternalLogger) Error(a ...interface{}) {
	common.PrintAndLog(fmt.Sprint(a))
}

func (logger *IrisExternalLogger) Warn(a ...interface{}) {
	common.PrintAndLog(fmt.Sprint(a))
}

func (logger *IrisExternalLogger) Info(a ...interface{}) {
	common.PrintAndLog(fmt.Sprint(a))
}

func (logger *IrisExternalLogger) Debug(a ...interface{}) {
	common.PrintAndLog(fmt.Sprint(a))
}
