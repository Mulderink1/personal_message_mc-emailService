package router

import (
	domain "github.com/Mulderink1/personal_message_mc-services-domain"
	"net/http"
)

func RunRoutes(log *domain.Logger) {
	log.Logger.Info("Register Routes")
	handlers := NewHandlers(log)

	http.HandleFunc("/message", handlers.MessageHandler)
}
