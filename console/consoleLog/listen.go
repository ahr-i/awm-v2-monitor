package consoleLog

import (
	"net"

	"github.com/ahr-i/awm-v2-monitor/setting"
	"github.com/ahr-i/awm-v2-monitor/src/logging"
	"github.com/ahr-i/awm-v2-monitor/src/logging/logDefault"
)

func listenLog() {
	receiverAddr, err := net.ResolveUDPAddr("udp", ":"+setting.Setting.ListenPort)
	if err != nil {
		logging.Logger.Warn(err)
	}

	receiverConn, err := net.ListenUDP("udp", receiverAddr)
	if err != nil {
		logging.Logger.Warn(err)
	}
	defer receiverConn.Close()

	buffer := make([]byte, setting.Setting.LogByte)
	for {
		n, _, err := receiverConn.ReadFromUDP(buffer)
		if err != nil {
			logging.Logger.Warn(err)
		}

		msg := string(buffer[:n])
		if authentication() {
			logDefault.Custom("[SERVICE]", msg)
		}
	}
}

func authentication() bool {
	return true
}
