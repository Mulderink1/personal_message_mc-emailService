package app

import (
	"fmt"
	"github.com/personal_message_mc-emailService/logger"
	"github.com/personal_message_mc-emailService/router"
	"net/http"
)

func Start() {
	log := logger.NewLogger()
	router.RunRoutes(log)

	PORT := "8080"

	log.Logger.Info(fmt.Sprintf("Listening on Port: %v", PORT))
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil)
}
