package webService

import (
	stdContext "context"
	"fmt"
	"github.com/Deansquirrel/go-tool"
	"github.com/Deansquirrel/goVerInfo/common"
	"github.com/Deansquirrel/goVerInfo/global"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var app *iris.Application

type WebService struct {
	Port int
}

func (ws *WebService) Start() {
	common.PrintAndLog("Service Start")
	app = iris.New()

	iris.RegisterOnInterrupt(ws.Stop)
	app.Use(recover.New())
	app.Use(logger.New())
	app.Logger().SetLevel("debug")

	if !global.SysConfig.Total.IsDebug {
		ws.setLogFile()
	}

	user := app.Party("/users", ws.myAuthMiddlewareHandler)

	user.Get("/{id:int}/profile", ws.userProfileHandler)

	user.Get("/inbox/{id:int}", ws.userMessageHandler)

	_ = app.Run(iris.Addr(":"+strconv.Itoa(ws.Port)), iris.WithoutInterruptHandler)
}

func (ws *WebService) Stop() {
	common.PrintAndLog("Service Stop")
	//关闭所有主机
	if app != nil {
		timeout := 5 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		_ = app.Shutdown(ctx)
	}
	return
}

func (ws *WebService) myAuthMiddlewareHandler(ctx iris.Context) {
	_, _ = ctx.WriteString("header\n")
	ctx.Next()
	_, _ = ctx.WriteString("\nfoot")
}
func (ws *WebService) userProfileHandler(ctx iris.Context) {
	id := ctx.Params().Get("id")
	fmt.Println("AAA - " + id)
	_, _ = ctx.WriteString(id)
}
func (ws *WebService) userMessageHandler(ctx iris.Context) {
	id := ctx.Params().Get("id")
	fmt.Println("BBB - " + id)
	_, _ = ctx.WriteString(id)
}

func (ws *WebService) setLogFile() {
	filePath := ws.getLogFileName()
	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		common.PrintAndLog("获取iris日志输出文件对象时遇到问题:" + err.Error())
		return
	}

	if app != nil {
		app.Logger().SetOutput(logFile)
		app.Logger().NewLine = true
	}
}

func (ws *WebService) getLogFileName() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	folderName := "logs"
	fileName := go_tool.GetDateStr(time.Now()) + ".log"
	if err != nil {
		dir = ""
	} else {
		dir = dir + "\\"
	}
	err = go_tool.CheckAndCreateFolder(dir + folderName)
	if err != nil {
		common.PrintAndLog("创建日志文件夹时遇到问题:" + err.Error())
		return fileName
	}
	fileName = dir + folderName + "\\" + "iris" + strconv.FormatInt(time.Now().Unix(), 10) + ".log"
	return fileName
}
