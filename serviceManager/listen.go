package serviceManager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"

	"github.com/ahr-i/awm-v2-monitor/setting"
	"github.com/ahr-i/awm-v2-monitor/src/logging"
	"github.com/ahr-i/awm-v2-monitor/src/logging/logDefault"
)

func ListenLog() {
	receiverAddr, err := net.ResolveUDPAddr("udp", ":"+setting.Setting.ListenPort)
	if err != nil {
		logging.Logger.Warn(err)
	}

	receiverConn, err := net.ListenUDP("udp", receiverAddr)
	if err != nil {
		logging.Logger.Error(err)
	}
	defer receiverConn.Close()

	buffer := make([]byte, setting.Setting.LogByte)
	for {
		n, address, err := receiverConn.ReadFromUDP(buffer)
		if err != nil {
			logging.Logger.Warn(err)
		}

		data := buffer[:n]
		service, msg, err := decodeJSON(data)
		if err != nil {
			logging.Logger.Warn(err)
		}

		ip, _, _ := net.SplitHostPort(address.String())
		if authentication(service, ip) {
			serviceFormat := fmt.Sprintf("[%s]", service)
			logDefault.Custom(serviceFormat, msg)
		} else {
			warringFormat := fmt.Sprintf("The message is not allowed. ! [%s] %s !", service, msg)
			logDefault.Warn(warringFormat)
		}
	}
}

func authentication(service string, ip string) bool {
	serviceInfo := Service[service]

	return serviceInfo.IP == ip
}

func decodeJSON(data []byte) (string, string, error) {
	var format listenFormat
	decoder := json.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&format)
	if err != nil {
		return "", "", err
	}

	return format.ServiceName, format.Message, nil
}
