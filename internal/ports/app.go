package ports

import (
	"context"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ichetiva/go-atm/internal/config"
	"github.com/ichetiva/go-atm/internal/controller"
	"github.com/ichetiva/go-atm/internal/router"
	"github.com/ichetiva/go-atm/internal/service"
	"github.com/ichetiva/go-atm/pkg/responder"
)

type app struct {
	log *slog.Logger
	cfg config.Config
}

func NewApp(log *slog.Logger, config config.Config) *app {
	return &app{
		log: log,
		cfg: config,
	}
}

func (a *app) Run() {
	a.log.Info("Bank initialization")
	bank := service.NewBank(a.log.With("service", "bank"))
	defer bank.Close()

	a.log.Info("Responder initialization")
	respond := responder.NewResponder(a.log.With("pkg", "responder"))

	a.log.Info("Controller initialization")
	controller := controller.NewController(bank, respond)

	a.log.Info("Router initialization")
	handler := router.NewRouter(a.log.With("layer", "router"), controller)

	a.log.Info("Run bank worker")
	go bank.RunWorkers(2)

	a.log.Info("Start listen address", "address", a.cfg.Addr)
	lis, err := net.Listen("tcp", a.cfg.Addr)
	if err != nil {
		log.Panic("Listen address error", err)
	}

	srv := http.Server{
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	if err := srv.Serve(lis); err != nil {
		log.Panic("Serve error", err)
	}

	<-sig

	srv.Shutdown(context.Background())
}
