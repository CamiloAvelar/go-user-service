package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/CamiloAvelar/go-user-service/infrastructure/config"
	"github.com/CamiloAvelar/go-user-service/infrastructure/database/mysql"
	serverHttp "github.com/CamiloAvelar/go-user-service/infrastructure/http"
	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
)

func main() {
	config := config.GetConfig()
	db := mysql.GetConnection(config)

	serverInjections := infrainterfaces.ServerInjections{
		Config: config,
		Db:     db,
	}

	s := serverHttp.GetServer(serverInjections)

	go func() {
		log.Println("Starting Server")
		if err := s.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Println("Server closed")
			} else {
				log.Printf("Server error: %v\n", err.Error())
			}
		}
	}()

	gracefulShutdown(s, db)
}

func gracefulShutdown(s serverHttp.Server, db *sql.DB) {
	var wg sync.WaitGroup

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	log.Println("Shutting down")

	if err := s.WaitShutdown(ctx); err != nil {
		log.Println(err.Error())
	}

	wg.Add(1)
	go func(swg *sync.WaitGroup) {
		defer swg.Done()
		if err := db.Close(); err != nil {
			log.Printf("Database close error: %v\n", err.Error())
		} else {
			log.Println("Database closed")
		}
	}(&wg)

	wg.Wait()

	log.Println("Bye")
	os.Exit(0)
}
