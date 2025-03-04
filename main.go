package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Udehlee/go-Ride/api/handlers"
	"github.com/Udehlee/go-Ride/api/routes"
	"github.com/Udehlee/go-Ride/db/db"
	"github.com/Udehlee/go-Ride/engine"
	"github.com/Udehlee/go-Ride/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	config, err := db.LoadConfig()
	if err != nil {
		log.Fatal("error loading config")
	}

	conn, err := db.InitDB(config)
	if err != nil {
		log.Fatal("error connecting to db")
	}

	workers := 5

	db := db.NewPgConn(conn.Conn)
	pq := engine.NewPriorityQueue()
	wp := engine.NewWorkerPool(workers, pq)

	log.Println("Starting the worker pool...")
	wp.Start()

	svc := service.NewService(&db, pq, wp)
	h := handlers.NewHandler(svc)

	routes.SetupRoutes(r, h)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("could not start server: %s\n", err)
		}
	}()

	// Wait for termination signal
	<-sigChan
	log.Println("Shutting down worker pool...")
	wp.Stop()
	log.Println("Worker pool shut down gracefully")
}
