# gowebarduino
Web remote console for arduino

# build
* go get -d -v
* go clean
* go build

# run
* go run main.go

# debug
* go get -u github.com/go-delve/delve/cmd/dlv
* dlv debug ./gowebarduino

# calls
* curl -X POST http://127.0.0.1:8000/api/v1/command -H 'Content-Type: application/json' -d '{"call": "3,2,21,github.com/wlanboy;"}'
* curl -X GET http://127.0.0.1:8000/api/v1/command 

# Command struct
*see great Arduino lib: https://github.com/thijse/Arduino-CmdMessenger 

# Arduino Sketch
* can be downloadde here: https://github.com/wlanboy/ArduinoSketches/tree/master/lcd_command
