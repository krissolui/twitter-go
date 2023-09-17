package main

import (
	"auth-service/internal/config"
	"auth-service/internal/routes"
	"auth-service/internal/server"
	"auth-service/internal/storage"
	"auth-service/internal/user"
	"log"
)

func main() {
	cfg := config.NewConfig()
	cfg.LoadEnv()
	store := storage.NewStorage()
	store.InitDB(cfg)

	userService := user.NewUserService(cfg, store)

	router := routes.NewRouter(userService)

	server := server.NewServer(cfg, router)
	if err := server.Start(); err != nil {
		log.Fatal("Failed to start server!")
	}
	defer server.Stop()
}
