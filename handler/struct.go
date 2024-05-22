package handler

import (
	"net/http"

	"github.com/unrolled/render"
)

var rend *render.Render = render.New()

type Handler struct {
	http.Handler
}

/* serviceRegisterHandler */
type registerRequest struct {
	ServiceName string `json:"service_name"`
	Port        string `json:"port"`
}
