![Go](https://github.com/wlanboy/gowebarduino/workflows/Go/badge.svg?branch=master)

# gowebarduino

HTTP remote console for Arduino. Exposes a REST API to send commands to and read responses from an Arduino connected via serial port (`/dev/ttyACM0`, 9600 baud).

## Requirements

- Go 1.26+
- Arduino connected via USB (`/dev/ttyACM0`)

## Dependencies

| Package | Purpose |
|---|---|
| `github.com/gorilla/mux` | HTTP router |
| `github.com/tarm/serial` | Serial port communication |

## Build

```bash
go mod download
go build -o gowebarduino .
```

## Run

```bash
go run main.go
```

The server listens on port `8000` by default. Override with the `PORT` environment variable:

```bash
PORT=9090 go run main.go
```

> If no Arduino is connected the server starts normally and returns an error on API calls.

## Docker

```bash
# Build binary first (Linux target for container)
GOOS=linux GOARCH=amd64 go build -o gowebarduino .

# Build and run image
docker build -t gowebarduino .
docker run --device=/dev/ttyACM0 -p 8000:8000 gowebarduino
```

## API

### POST /api/v1/command — send a command

```bash
curl -X POST http://127.0.0.1:8000/api/v1/command \
  -H 'Content-Type: application/json' \
  -d '{"call": "3,2,21,github.com/wlanboy;"}'
```

Response `201 Created`:
```json
{ "call": "3,2,21,github.com/wlanboy;", "result": "" }
```

### GET /api/v1/command — read last response

```bash
curl -X GET http://127.0.0.1:8000/api/v1/command
```

Response `200 OK`:
```json
{ "call": "", "result": "OK" }
```

## Debug

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
dlv debug .
```

## Command format

Commands follow the [Arduino-CmdMessenger](https://github.com/thijse/Arduino-CmdMessenger) protocol:

```
<commandId>,<arg1>,<arg2>,...,<argN>;
```

## Arduino Sketch

Compatible sketch available at: https://github.com/wlanboy/ArduinoSketches/tree/master/lcd_command
