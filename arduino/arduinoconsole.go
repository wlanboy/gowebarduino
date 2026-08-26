package arduino

import (
	"io"
	"sync"
	"time"

	serial "github.com/tarm/serial"
)

/*Arduino struct*/
type Arduino struct {
	Name    string
	Console io.ReadWriteCloser
	mu      sync.Mutex
}

/*CreateConsole to an Arduino on Device ttyACM0*/
func CreateConsole() (*Arduino, error) {
	arduino := &Arduino{}

	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600, ReadTimeout: time.Second * 5}
	s, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}
	time.Sleep(1 * time.Second)

	arduino.Console = s
	arduino.Name = "Arudino Uno"

	return arduino, nil
}

/*WriteConsole on exportet Console*/
func WriteConsole(command string, arduino *Arduino) error {
	var consoleerror error

	arduino.mu.Lock()
	defer arduino.mu.Unlock()

	_, err := arduino.Console.Write([]byte(command))
	if err != nil {
		consoleerror = err
	}
	return consoleerror
}

/*ReadConsole on exportet Console*/
func ReadConsole(arduino *Arduino) (string, error) {
	var returnstring string
	var consoleerror error

	arduino.mu.Lock()
	defer arduino.mu.Unlock()

	buf := make([]byte, 128)
	number, err := arduino.Console.Read(buf)
	if err != nil {
		consoleerror = err
	}
	returnstring = string(buf[:number])
	return returnstring, consoleerror
}
