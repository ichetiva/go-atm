package router

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/ichetiva/go-atm/docs"
	"github.com/ichetiva/go-atm/internal/controller"
	slogchi "github.com/samber/slog-chi"
	httpHandler "github.com/swaggo/http-swagger/v2"
)

func NewRouter(log *slog.Logger, c *controller.Controller) http.Handler {
	r := chi.NewRouter()

	// middlewares
	r.Use(slogchi.New(log))
	r.Use(middleware.Recoverer)

	// account routes
	r.Post("/accounts", c.CreateAccount)
	r.Post("/accounts/{id}/deposit", c.Deposit)
	r.Post("/accounts/{id}/withdraw", c.Withdraw)
	r.Get("/accounts/{id}/balance", c.GetBalance)

	// swagger route
	r.Get("/swagger/*", httpHandler.Handler(httpHandler.URL("http://localhost:8000/swagger/doc.json")))

	return r
}
