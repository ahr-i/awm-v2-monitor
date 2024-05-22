package serviceManager

import "sync"

var Service map[string]ServiceInfo

type ServiceInfo struct {
	IP      string
	Address []string
}

var mutex sync.Mutex
