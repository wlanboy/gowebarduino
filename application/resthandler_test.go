package application

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	arduino "github.com/wlanboy/gowebarduino/arduino"
	model "github.com/wlanboy/gowebarduino/model"
)

/*fakeConsole is an in-memory io.ReadWriteCloser standing in for the serial port*/
type fakeConsole struct {
	written  []byte
	toRead   []byte
	writeErr error
	readErr  error
}

func (f *fakeConsole) Write(p []byte) (int, error) {
	if f.writeErr != nil {
		return 0, f.writeErr
	}
	f.written = append(f.written, p...)
	return len(p), nil
}

func (f *fakeConsole) Read(p []byte) (int, error) {
	if f.readErr != nil {
		return 0, f.readErr
	}
	n := copy(p, f.toRead)
	return n, nil
}

func (f *fakeConsole) Close() error { return nil }

func TestPostCreateWithoutConsole(t *testing.T) {
	goservice := &GoService{}
	body := bytes.NewBufferString(`{"call":"on"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/command", body)
	w := httptest.NewRecorder()

	goservice.PostCreate(w, req)

	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusServiceUnavailable)
	}
}

func TestPostCreateInvalidJSON(t *testing.T) {
	goservice := &GoService{}
	body := bytes.NewBufferString(`not json`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/command", body)
	w := httptest.NewRecorder()

	goservice.PostCreate(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

func TestPostCreateEmptyCall(t *testing.T) {
	goservice := &GoService{Console: &arduino.Arduino{Console: &fakeConsole{}}}
	body := bytes.NewBufferString(`{"call":""}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/command", body)
	w := httptest.NewRecorder()

	goservice.PostCreate(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

func TestPostCreateSuccess(t *testing.T) {
	fake := &fakeConsole{}
	goservice := &GoService{Console: &arduino.Arduino{Console: fake}}
	body := bytes.NewBufferString(`{"call":"on"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/command", body)
	w := httptest.NewRecorder()

	goservice.PostCreate(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusCreated)
	}
	if string(fake.written) != "on" {
		t.Errorf("written to console = %q, want %q", fake.written, "on")
	}
}

func TestPostCreateConsoleError(t *testing.T) {
	fake := &fakeConsole{writeErr: errors.New("boom")}
	goservice := &GoService{Console: &arduino.Arduino{Console: fake}}
	body := bytes.NewBufferString(`{"call":"on"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/command", body)
	w := httptest.NewRecorder()

	goservice.PostCreate(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusInternalServerError)
	}
}

func TestGetWithoutConsole(t *testing.T) {
	goservice := &GoService{}
	req := httptest.NewRequest(http.MethodGet, "/api/v1/command", nil)
	w := httptest.NewRecorder()

	goservice.Get(w, req)

	if w.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusServiceUnavailable)
	}
}

func TestGetSuccess(t *testing.T) {
	fake := &fakeConsole{toRead: []byte("hello")}
	goservice := &GoService{Console: &arduino.Arduino{Console: fake}}
	req := httptest.NewRequest(http.MethodGet, "/api/v1/command", nil)
	w := httptest.NewRecorder()

	goservice.Get(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusOK)
	}
	var got model.Command
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatalf("response is not valid JSON: %v", err)
	}
	if got.Result != "hello" {
		t.Errorf("result = %q, want %q", got.Result, "hello")
	}
}
