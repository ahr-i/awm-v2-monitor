package main

import (
	"net/http"

	"github.com/ahr-i/awm-v2-monitor/console"
	"github.com/ahr-i/awm-v2-monitor/handler"
	"github.com/ahr-i/awm-v2-monitor/serviceManager"
	"github.com/ahr-i/awm-v2-monitor/setting"
	"github.com/ahr-i/awm-v2-monitor/src/corsController"
	"github.com/ahr-i/awm-v2-monitor/src/logging"
	"github.com/urfave/negroni"
)

func initialization() {
	setting.Init()
	logging.Init()
	console.Init()
	serviceManager.Init()
}

func startServer() {
	mux := handler.CreateHandler()
	handler := negroni.Classic()
	defer mux.Close()

	handler.Use(corsController.SetCors("*", "GET, POST, PUT, DELETE", "*", true))
	handler.UseHandler(mux)

	logging.Logger.Info("HTTP server start.")
	http.ListenAndServe(":"+setting.Setting.ServerPort, handler)
}

func main() {
	initialization()
	startServer()
}
