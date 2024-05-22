package handler

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/ahr-i/awm-v2-monitor/serviceManager"
	"github.com/ahr-i/awm-v2-monitor/src/logging/logDefault"
)

func (h *Handler) serviceRegisterHandler(w http.ResponseWriter, r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	var request registerRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		rend.JSON(w, http.StatusBadRequest, nil)
		return
	}

	serviceManager.Register(request.ServiceName, ip, request.Port)

	logDefault.Error(ip + " " + request.ServiceName)
	rend.JSON(w, http.StatusOK, nil)
}
