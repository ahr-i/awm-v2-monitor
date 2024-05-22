package setting

var Setting settingStruct

type settingStruct struct {
	ServerPort  string            `json:"server_port"`
	ListenPort  string            `json:"listen_port"`
	LogByte     int               `json:"log_byte"`
	Service     map[string]string `json:"service"`
	HealthCheck healthCheck       `json:"health_check"`
}

type healthCheck struct {
	Use     bool `json:"Use"`
	Time    int  `json:"time"`
	Timeout int  `json:"timeout"`
}
