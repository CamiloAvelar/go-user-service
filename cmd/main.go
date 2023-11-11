package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/CamiloAvelar/go-user-service/infrastructure/config"
	"github.com/CamiloAvelar/go-user-service/infrastructure/database/mysql"
	"github.com/CamiloAvelar/go-user-service/infrastructure/http"
	infrainterfaces "github.com/CamiloAvelar/go-user-service/infrastructure/interfaces"
)

func main() {
	config := config.GetConfig()
	db := mysql.GetConnection(config)

	httpInjections := infrainterfaces.HttpServerInjections{
		Config: config,
		Db:     db,
	}

	s := http.Get(httpInjections)
	server := s.Server

	go func() {
		log.Println("Starting Server")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	gracefulShutdown(s, db)
}

func gracefulShutdown(s http.Server, db *sql.DB) {
	var wg sync.WaitGroup

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	log.Println("Shutting down")

	s.WaitShutdown(ctx)

	wg.Add(1)
	go func() {
		defer wg.Done()
		db.Close()
	}()

	wg.Wait()

	log.Println("Bye")
	os.Exit(0)
}
