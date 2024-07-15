package responder

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Responder interface {
	Ok(w http.ResponseWriter, data interface{})
	ErrorBadRequest(w http.ResponseWriter, err error)
	ErrorInternal(w http.ResponseWriter, err error)
}

type responder struct {
	log *slog.Logger
}

type Response struct {
	Ok      bool        `json:"ok"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponder(log *slog.Logger) Responder {
	return &responder{
		log: log,
	}
}

func (r *responder) Ok(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response{
		Ok:      true,
		Message: "",
		Data:    data,
	}); err != nil {
		r.log.Error("Encode OK response", "error", err)
	}
}

func (r *responder) ErrorBadRequest(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(Response{
		Ok:      false,
		Message: err.Error(),
		Data:    nil,
	}); err != nil {
		r.log.Error("Encode ERROR response", "error", err)
	}
}

func (r *responder) ErrorInternal(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{
		Ok:      false,
		Message: err.Error(),
		Data:    nil,
	}); err != nil {
		r.log.Error("Encode ERROR response", "error", err)
	}
}
