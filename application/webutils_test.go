package application

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	model "github.com/wlanboy/gowebarduino/model"
)

func TestWriteJSONErrorResponseEscapesMessage(t *testing.T) {
	w := httptest.NewRecorder()
	WriteJSONErrorResponse(w, `message with "quotes" and`+"\n newline", 400)

	var body map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("response is not valid JSON: %v (body: %s)", err, w.Body.String())
	}
	if body["error"] == "" {
		t.Errorf("expected error field to be set, got %q", body["error"])
	}
	if w.Code != 400 {
		t.Errorf("status = %d, want 400", w.Code)
	}
}

func TestWriteJSONResponse(t *testing.T) {
	w := httptest.NewRecorder()
	command := model.Command{Call: "on", Result: "ok"}
	WriteJSONResponse(w, &command, 201)

	var got model.Command
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatalf("response is not valid JSON: %v", err)
	}
	if got != command {
		t.Errorf("got %+v, want %+v", got, command)
	}
	if w.Code != 201 {
		t.Errorf("status = %d, want 201", w.Code)
	}
}
