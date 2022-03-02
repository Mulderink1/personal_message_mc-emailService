package router

import (
	domain "github.com/Mulderink1/personal_message_mc-services-domain"
	"github.com/personal_message_mc-emailService/email"
	"net/http"
)

type Handlers struct {
	Log *domain.Logger
}

func NewHandlers(log *domain.Logger) *Handlers {
	return &Handlers{
		Log: log,
	}
}

func (h *Handlers) MessageHandler(w http.ResponseWriter, r *http.Request) {
	em := r.URL.Query().Get("email")

	h.Log.Logger.Info("MessageHandler received Get")
	email.SmsSender(h.Log, em)
}
