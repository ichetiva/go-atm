package controller

import (
	"github.com/ichetiva/go-atm/internal/service"
	"github.com/ichetiva/go-atm/pkg/responder"
)

type Controller struct {
	bank *service.Bank
	responder.Responder
}

func NewController(bank *service.Bank, respond responder.Responder) *Controller {
	return &Controller{
		bank:      bank,
		Responder: respond,
	}
}
