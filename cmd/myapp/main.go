package main

import (
	"fmt"
	"log"

	"myapp/internal/cache"
	"myapp/internal/config"
	"myapp/internal/db"
	"myapp/internal/domain/todo"
	"myapp/internal/server"
)

func main() {
	cfg := config.LoadConfig()

	// DB
	dbs, err := db.NewDBs(cfg.Databases)
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}
	defer dbs.MySQL.Close()

	// Cache
	cacheClient := cache.NewRedis(cfg.Cache.Address, cfg.Cache.Password, cfg.Cache.DB)

	// Todo wiring
	todoRepo := todo.NewRepository(dbs.MySQL, cacheClient)
	todoService := todo.NewService(todoRepo)
	todoHandler := todo.NewHandler(todoService)

	// Server
	r := server.NewServer(todoHandler)
	addr := fmt.Sprintf(":%d", cfg.App.Port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
