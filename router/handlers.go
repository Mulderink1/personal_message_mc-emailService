package router

import (
	"encoding/json"
	"fmt"
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
	if r.Method == http.MethodPost {
		m, err := decodeMessage(r, h)
		if err != nil {
			return
		}
		err = email.SmsSender(h.Log, m)
		if err != nil {
			h.Log.Logger.Error(fmt.Sprintf("error: %s: %v", err.Error(), http.StatusBadRequest))
			writeResponse(w, 500, err.Error())
		} else {
			writeResponse(w, 200, "Success")
		}
	}
}

func decodeMessage(r *http.Request, h *Handlers) (domain.Message, error) {
	h.Log.Logger.Info("MessageHandler received Get")

	m := domain.Message{}
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		h.Log.Logger.Error(fmt.Sprintf("error: %s: %v", err.Error(), http.StatusBadRequest))
		return domain.Message{}, err
	}

	return m, nil
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
