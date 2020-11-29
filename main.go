package main

import (
	"github.com/airbrake/gobrake/v5"
	"net/http"
)

var _ = gobrake.NewNotifierWithOptions(&gobrake.NotifierOptions{
	ProjectId:   311603,
	ProjectKey:  "4f663f741aa2a901e6619f2e8a9d93b1",
	Environment: "production",
})

func main() {
	var server http.Server
	//defer airbrake.Close()
	//defer airbrake.NotifyOnPanic()
	server.Addr = ":8080"
	server.Handler = Router()
	server.ListenAndServe()
}
