package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "minecraft/docs"
	"minecraft/internal/book"
	"minecraft/internal/config"
	"minecraft/internal/db"
	"minecraft/internal/swagger"
)

// @title			MinecraftTest
// @version		1.0
// @description	This is a dummy.
func main() {
	cfg := config.Load()

	// Подключаемся к БД
	database, err := db.New(cfg)
	if err != nil {
		log.Fatal("Ошибка подключения к базе: ", err)
	}
	defer database.Close()

	// Initialize Service, Controller and Router
	bService := book.NewService()
	bController := book.NewHandler(bService)
	book.SetupBooksRouter(bController)

	swagger.SetupSwaggerRouter()

	server := &http.Server{
		Addr: ":8080",
	}

	log.Printf("🚀 Приложение запущено. Нажмите Ctrl+C для graceful shutdown.")

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	server.Shutdown(ctx)
	database.Close()

	log.Println("Приложение успешно остановлено.")
}
