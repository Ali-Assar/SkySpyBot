package main

import (
	"log"
	"net/http"
	"os"

	database "github.com/Ali-Assar/SkySpyBot/db"
	"github.com/Ali-Assar/SkySpyBot/handler"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	handler.OWMApiKey = os.Getenv("OWM_API_KEY")
	handler.TelegramApikey = os.Getenv("TELEGRAM_BOT_TOKEN")
	redisAddress := os.Getenv("REDIS_ADDRESS")

	log.Printf("handler.OWMApiKey:%s handler.TelegramApikey:%s redisAddress:%s\n", handler.OWMApiKey, handler.TelegramApikey, redisAddress)

	redisClient, cancel, err := database.NewRedisClient(redisAddress)
	if err != nil {
		log.Fatalf("Failed to create Redis client: %v", err)
	}
	defer cancel()

	handler.RedisClient = redisClient // set the RedisClient in the handler package

	log.Println("running")
	http.ListenAndServe(":5000", http.HandlerFunc(handler.Handler))
}
