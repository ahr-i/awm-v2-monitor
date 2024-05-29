package serviceManager

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ahr-i/awm-v2-monitor/src/logging"
)

func healthCheck(checkTime int, timeout int) {
	for {
		mutex.Lock()
		for serviceKey, serviceInfo := range Service {
			var newAddress []string

			for _, address := range serviceInfo.Address {
				url := fmt.Sprintf("http://%s/ping", address)

				result := checkPingRequest(url, timeout)
				if result {
					newAddress = append(newAddress, address)
				} else {
					msg := fmt.Sprintf("Delete service address: %s", address)
					logging.Logger.Warn(msg)
				}
			}
			serviceInfo.Address = newAddress
			Service[serviceKey] = serviceInfo
		}
		mutex.Unlock()

		time.Sleep(time.Duration(checkTime) * time.Second)
	}
}

func checkPingRequest(url string, timeout int) bool {
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	_, err := client.Get(url)

	return err == nil
}
