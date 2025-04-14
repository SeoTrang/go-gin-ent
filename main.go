package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/seotrang/go-ent/ent"
	"github.com/seotrang/go-ent/routes"
)

func main() {
	// user:password@tcp(host:port)/dbname
	client, err := ent.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_ent_db?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	// Migrate schema vào DB thật
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	r := gin.Default()
	routes.SetupRoutes(r, client)

	log.Println("Server is running on http://localhost:8080")
	r.Run(":8080")
}
