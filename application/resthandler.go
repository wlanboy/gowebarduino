package application

import (
	"encoding/json"
	"log"
	"net/http"

	arduino "github.com/wlanboy/gowebarduino/arduino"
	model "github.com/wlanboy/gowebarduino/model"
)

/*PostCreate POST method*/
func (goservice *GoService) PostCreate(w http.ResponseWriter, r *http.Request) {

	command := model.Command{}

	err := json.NewDecoder(r.Body).Decode(&command)
	if err != nil {
		log.Println("Cannot parse JSON")
		log.Println(err)
		WriteJSONErrorResponse(w, "Cannot parse JSON", http.StatusBadRequest)
		return
	}

	if message, valid := command.Validate(); !valid {
		WriteJSONErrorResponse(w, message, http.StatusBadRequest)
		return
	}

	if goservice.Console == nil {
		WriteJSONErrorResponse(w, "Arduino console is not available", http.StatusServiceUnavailable)
		return
	}

	consoleerror := arduino.WriteConsole(command.Call, goservice.Console)
	if consoleerror != nil {
		log.Println("Console error")
		log.Println(consoleerror)
		WriteJSONErrorResponse(w, consoleerror.Error(), http.StatusInternalServerError)
		return
	}
	WriteJSONResponse(w, &command, http.StatusCreated)
}

/*Get method*/
func (goservice *GoService) Get(w http.ResponseWriter, r *http.Request) {
	command := model.Command{}

	if goservice.Console == nil {
		WriteJSONErrorResponse(w, "Arduino console is not available", http.StatusServiceUnavailable)
		return
	}

	resp, consoleerror := arduino.ReadConsole(goservice.Console)
	if consoleerror != nil {
		log.Println("Console error")
		log.Println(consoleerror)
		WriteJSONErrorResponse(w, consoleerror.Error(), http.StatusInternalServerError)
		return
	}
	command.Result = resp
	WriteJSONResponse(w, &command, http.StatusOK)
}
