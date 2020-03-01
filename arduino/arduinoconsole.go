package arduino

import (
	"io"
	"log"
	"time"

	serial "github.com/tarm/serial"
)

/*Arduino struct*/
type Arduino struct {
	Name    string
	Console io.ReadWriteCloser
}

/*CreateConsole to an Arduino on Device ttyACM0*/
func CreateConsole() Arduino {
	var arduino Arduino

	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 9600, ReadTimeout: time.Second * 5}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(1 * time.Second)

	arduino.Console = s
	arduino.Name = "Arudino Uno"

	return arduino
}

/*WriteConsole on exportet Console*/
func WriteConsole(command string, arduino *Arduino) error {
	var consoleerror error

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

	buf := make([]byte, 128)
	number, err := arduino.Console.Read(buf)
	if err != nil {
		consoleerror = err
	}
	returnstring = string(buf[:number])
	return returnstring, consoleerror
}
