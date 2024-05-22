package logging

import "github.com/ahr-i/awm-v2-monitor/src/logging/loggingIPFS"

func Init() {
	Logger = newLogger()

	Logger.Init()
}

func newLogger() Logging {
	return loggingIPFS.NewLogger()
}
