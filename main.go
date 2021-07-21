package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/danishk121/go-frame/handler"
	"github.com/danishk121/go-frame/service"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func main() {

	l := log.New(os.Stdout, "LO/LI-api ", log.LstdFlags)
	d, _ := ConfigureDatabase()
	s := service.NewLOService(l, d)

	// create the handlers
	lh := handler.NewLO(l, s)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	LORouter := sm.PathPrefix("/lo/").Subrouter()
	LIRouter := sm.PathPrefix("/li/").Subrouter()
	InitLORoutes(LORouter, lh)
	InitLIRoutes(LIRouter)

	StartAndGracefullShutdown(l, sm)
}

func InitLORoutes(sm *mux.Router, lh *handler.LO) {

	const V1API = "/api/v1"
	x := sm.Methods(http.MethodPost).Subrouter()
	x.Use(lh.MiddlewareValidateProduct)
	x.HandleFunc(V1API, lh.AddLO)

	sm.Methods(http.MethodGet).Subrouter().HandleFunc(V1API, lh.GetLO)

}

func InitLIRoutes(sm *mux.Router) {

}

func ConfigureDatabase() (*gorm.DB, error) {

	var appConfig map[string]string
	appConfig, confErr := godotenv.Read()

	if confErr != nil {
		log.Fatal("Error reading .env file")
		return nil, confErr
	}

	// Ex: user:password@tcp(host:port)/dbname
	mysqlCredentials := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		appConfig["MYSQL_USER"],
		appConfig["MYSQL_PASSWORD"],
		appConfig["MYSQL_PROTOCOL"],
		appConfig["MYSQL_HOST"],
		appConfig["MYSQL_PORT"],
		appConfig["MYSQL_DBNAME"],
	)

	database, confErr := gorm.Open(mysql.Open(mysqlCredentials), &gorm.Config{})

	if confErr != nil {
		log.Printf("Error starting server: %s\n", confErr)
		os.Exit(1)
	}
	return database, confErr

}

func StartAndGracefullShutdown(l *log.Logger, sm *mux.Router) {
	// create a new server
	s := http.Server{
		Addr:         ":8081",           // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}
	// start the server
	go func() {
		l.Println("Starting server on port 8081")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		l.Fatal(err)
	}
	error := s.Shutdown(ctx)
	if err != nil {
		l.Fatal(error)
	}

}
