package console

import "github.com/ahr-i/awm-v2-monitor/console/consoleLog"

func Init() {
	go consoleLog.Init()
}
