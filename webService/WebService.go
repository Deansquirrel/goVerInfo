package webService

import (
	stdContext "context"
	"fmt"
	"github.com/Deansquirrel/goVerInfo/common"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"strconv"
	"time"
)

var app *iris.Application

type WebService struct {
	Port int
}

func (ws *WebService) Start() {
	common.PrintAndLog("Service Starting")
	app = iris.New()
	app.Logger().SetLevel("debug")

	iris.RegisterOnInterrupt(ws.Stop)
	app.Use(recover.New())
	app.Use(logger.New())

	user := app.Party("/users", ws.myAuthMiddlewareHandler)

	user.Get("/{id:int}/profile", ws.userProfileHandler)

	user.Get("/inbox/{id:int}", ws.userMessageHandler)

	_ = app.Run(iris.Addr(":"+strconv.Itoa(ws.Port)), iris.WithoutInterruptHandler)
	//if err != nil {
	//	common.PrintAndLog("服务启动时遇到错误:" + err.Error())
	//} else {
	//	common.PrintAndLog("Service Started")
	//}
}

func (ws *WebService) Stop() {
	//关闭所有主机
	if app != nil {
		timeout := 5 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		err := app.Shutdown(ctx)
		if err != nil {
			common.PrintAndLog("ws关闭时遇到错误:" + err.Error())
		}
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
