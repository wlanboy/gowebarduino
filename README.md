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

# dockerize (docker image size is 9.89MB)
* GOOS=linux GOARCH=386 go build (386 needed for busybox)
* GOOS=linux GOARCH=arm GOARM=6 go build (Raspberry Pi build)
* GOOS=linux GOARCH=arm64 go build (Odroid C2 build)
* docker build -t gowebarduino .

# run docker container
*docker run -d -p 8000:8000 gowebarduino

# calls
* curl -X POST http://127.0.0.1:8000/api/v1/command -H 'Content-Type: application/json' -d '{"call": "3,2,21,github.com/wlanboy;"}'
* curl -X GET http://127.0.0.1:8000/api/v1/command 

#Command struct
*see great Arduino lib: https://github.com/thijse/Arduino-CmdMessenger 
