package main

import (
	"log"
	"net/http"

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
	wp.Start()

	svc := service.NewService(&db, pq, wp)
	h := handlers.NewHandler(svc)

	routes.SetupRoutes(r, h)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
