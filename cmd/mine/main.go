package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"minecraft/internal/config"
	"minecraft/internal/db"
)

func main() {
	cfg := config.Load()

	// Подключаемся к БД
	database, err := db.New(cfg)
	if err != nil {
		log.Fatal("Ошибка подключения к базе:", err)
	}
	defer database.Close()

	// Здесь позже будет инициализация роутера, сервер и т.д.

	log.Printf("🚀 Приложение запущено. Нажмите Ctrl+C для graceful shutdown.")

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Получен сигнал завершения, останавливаем приложение...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := database.Close(); err != nil {
		log.Printf("Ошибка при закрытии БД: %v", err)
	}

	log.Println("Приложение успешно остановлено.")
}
