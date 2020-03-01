package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	arduino "../arduino"
	"github.com/gorilla/mux"
)

/*GoService containing router and database*/
type GoService struct {
	Router  *mux.Router
	Server  *http.Server
	Console *arduino.Arduino
}

/*Run app and initialize db connection and http server*/
func (goservice *GoService) Run() {

	//load http server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println(port)

	goservice.Server = &http.Server{
		Handler:      goservice.Router,
		Addr:         fmt.Sprintf(":%s", port),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Println("Starting http server...")
		if err := goservice.Server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

}

/*WaitForShutdown application server*/
func (goservice *GoService) WaitForShutdown() {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	goservice.Server.Shutdown(ctx)

	log.Println("Shutting down http server.")
	os.Exit(0)
}
