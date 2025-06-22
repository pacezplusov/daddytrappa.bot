package main

import (
	"Ruslan/internal/bot"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mymmrac/telego"
)

func main() {
	// Получаем токен из переменной окружения
	botToken := os.Getenv("botToken")
	if botToken == "" {
		log.Fatal("Bot token is required")
	}

	// Создаём нового бота
	b, err := telego.NewBot(botToken)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// Запускаем бота в отдельной горутине
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered panic in bot goroutine: %v", r)
			}
		}()
		log.Println("Starting bot...")
		if err := bot.StartBot(b); err != nil {
			// Не используем log.Fatal, чтобы не убить main-горутина
			log.Printf("Bot stopped with error: %v", err)
		}
	}()

	// Инициализация Gin-сервера
	r := gin.Default()

	// Пример API-эндпоинта для вывода суммы долга
	r.GET("/shlyapster/dolg", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Сумма долга": 91000,
		})
	})

	// Запускаем HTTP-сервер в отдельной горутине
	go func() {
		if err := r.Run(); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	// Блокируем main, чтобы программа не завершалась
	select {}
}
