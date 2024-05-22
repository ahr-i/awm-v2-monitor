package serviceManager

import (
	"fmt"

	"github.com/ahr-i/awm-v2-monitor/src/logging"
)

func Register(service string, ip string, port string) {
	mutex.Lock()
	defer mutex.Unlock()

	address := fmt.Sprintf("%s:%s", ip, port)
	serviceInfo, exist := Service[service]
	if !exist {
		rejectLog(service, address)

		return
	}

	if serviceInfo.IP != ip {
		rejectLog(service, address)

		return
	}

	Service[service] = ServiceInfo{
		IP:      serviceInfo.IP,
		Address: append(serviceInfo.Address, address),
	}
	logging.Logger.Info("Service register: " + service + " / Address: " + address)
}

func rejectLog(service string, address string) {
	msg := fmt.Sprintf("The [%s] service does not exist: Reject.", service)

	logging.Logger.Info(msg)
	logging.Logger.Info("Reject Address: " + address)
}
