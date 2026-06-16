package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"minecraft/internal/config"
	"minecraft/internal/db"
	bookController "minecraft/internal/server/controllers"
	bookRouter "minecraft/internal/server/routers"
	bookService "minecraft/internal/server/services"
)

func main() {
	cfg := config.Load()

	// Подключаемся к БД
	database, err := db.New(cfg)
	if err != nil {
		log.Fatal("Ошибка подключения к базе:", err)
	}
	defer database.Close()

	// Initialize Service, Controller and Router
	bService := bookService.NewBookService()
	bController := bookController.NewBookController(bService)
	bookRouter.SetupBooksRouter(bController)

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
