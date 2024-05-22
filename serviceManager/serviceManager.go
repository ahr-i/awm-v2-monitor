package serviceManager

import (
	"github.com/ahr-i/awm-v2-monitor/setting"
	"github.com/ahr-i/awm-v2-monitor/src/logging"
)

func Init() {
	setService()

	if setting.Setting.HealthCheck.Use {
		logging.Logger.Info("Start the Service Health Check.")
		go healthCheck(setting.Setting.HealthCheck.Time, setting.Setting.HealthCheck.Timeout)
	}
}

func setService() {
	Service = make(map[string]ServiceInfo)

	for service, ip := range setting.Setting.Service {
		var serviceInfo = ServiceInfo{
			IP:      ip,
			Address: nil,
		}

		Service[service] = serviceInfo
	}
	logging.Logger.Info("Generate a list of services.")
	logging.Logger.Info(Service)
}
