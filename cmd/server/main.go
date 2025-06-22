package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"Soraka/internal/db"
	"Soraka/internal/router"
)

func main() {
	if err := db.Init(); err != nil {
		log.Fatalf("init db: %v", err)
	}
	r := gin.Default()
	router.Register(r)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
