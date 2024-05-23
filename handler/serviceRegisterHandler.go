package handler

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/ahr-i/awm-v2-monitor/serviceManager"
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

	rend.JSON(w, http.StatusOK, nil)
}
