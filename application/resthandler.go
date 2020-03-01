package application

import (
	"encoding/json"
	"log"
	"net/http"
	arduino "../arduino"
	model "../model"
)

/*PostCreate POST method*/
func (goservice *GoService) PostCreate(w http.ResponseWriter, r *http.Request) {

	command := model.Command{}

	err := json.NewDecoder(r.Body).Decode(&command)
	if err != nil {
		log.Println("Cannot parse JSON")
		log.Println(err)
		WriteJSONErrorResponse(w, "Cannot parse JSON", http.StatusBadRequest)
	} else {
		consoleerror := arduino.WriteConsole(command.Call, goservice.Console)
		if consoleerror != nil {
			log.Println("Console error")
			log.Println(consoleerror)
			WriteJSONErrorResponse(w, consoleerror.Error(), http.StatusInternalServerError)
		} else {
			WriteJSONResponse(w, &command, http.StatusCreated)
		}
	}
}

/*Get method*/
func (goservice *GoService) Get(w http.ResponseWriter, r *http.Request) {
	command := model.Command{}

	resp,consoleerror := arduino.ReadConsole(goservice.Console)
	if consoleerror != nil {
		log.Println("Console error")
		log.Println(consoleerror)
		WriteJSONErrorResponse(w, consoleerror.Error(), http.StatusInternalServerError)
	} else {
		command.Result = resp
		WriteJSONResponse(w, &command, http.StatusCreated)
	}
}
