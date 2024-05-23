package consoleLog

import "github.com/ahr-i/awm-v2-monitor/serviceManager"

func Init() {
	serviceManager.ListenLog()
}
