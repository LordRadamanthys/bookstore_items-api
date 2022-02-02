package controllers

import "net/http"

const pong = "pong!"

var (
	PinController pingControllerInterface = &pingController{}
)

type pingController struct{}
type pingControllerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

func (p *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
}
