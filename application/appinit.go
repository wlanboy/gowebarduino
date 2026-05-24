package application

import (
	"log"

	arduino "github.com/wlanboy/gowebarduino/arduino"
	"github.com/gorilla/mux"
)

/*Initialize app router and configuration*/
func (goservice *GoService) Initialize() {
	goservice.Router = mux.NewRouter()

	goservice.Router.HandleFunc("/api/v1/command", goservice.PostCreate).Methods("POST")
	goservice.Router.HandleFunc("/api/v1/command", goservice.Get).Methods("GET")

	console, err := arduino.CreateConsole()
	if err != nil {
		log.Printf("Warning: could not open Arduino serial port: %v", err)
		return
	}
	goservice.Console = &console
}

