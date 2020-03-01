package application

import (
	arduino "../arduino"
	"github.com/gorilla/mux"
)

/*Initialize app router and configuration*/
func (goservice *GoService) Initialize() {
	goservice.Router = mux.NewRouter()

	goservice.Router.HandleFunc("/api/v1/command", goservice.PostCreate).Methods("POST")
	goservice.Router.HandleFunc("/api/v1/command", goservice.Get).Methods("Get")

	var console arduino.Arduino = arduino.CreateConsole();
	goservice.Console = &console
}

