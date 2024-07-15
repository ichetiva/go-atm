package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// @summary CreateAccount
// @tags account
// @accept json
// @produce json
// @success 200 {object} responder.Response "Success"
// @router /accounts [post]
func (a *Controller) CreateAccount(w http.ResponseWriter, r *http.Request) {
	id := a.bank.CreateAccount()
	a.Ok(w, createAccountResponse{ID: id})
}

// @summary Deposit cash to account
// @tags account
// @accept json
// @produce json
// @param id path string true "a202482a-bf68-4e17-a4c4-5a268b2e15d6"
// @param amount body float32 true "Deposit amount"
// @success 200 {object} responder.Response "Success"
// @failure 400 {object} responder.Response "Bad request"
// @failure 500 {object} responder.Response "Internal error"
// @router /accounts/{id}/deposit [post]
func (a *Controller) Deposit(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		a.ErrorBadRequest(w, fmt.Errorf("id is empty"))
		return
	}

	var req operationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		a.ErrorBadRequest(w, err)
		return
	}

	account := a.bank.GetAccount(id)
	if err := account.Deposit(req.Amount); err != nil {
		a.ErrorInternal(w, err)
		return
	}

	a.Ok(w, nil)
}

// @summary Withdraw cash from account
// @tags account
// @accept json
// @produce json
// @param id path string true "a202482a-bf68-4e17-a4c4-5a268b2e15d6"
// @param amount body float32 true "Withdraw amount"
// @success 200 {object} responder.Response "Success"
// @failure 400 {object} responder.Response "Bad request"
// @failure 500 {object} responder.Response "Internal error"
// @router /accounts/{id}/withdraw [post]
func (a *Controller) Withdraw(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		a.ErrorBadRequest(w, fmt.Errorf("id is empty"))
		return
	}

	var req operationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		a.ErrorBadRequest(w, err)
		return
	}

	account := a.bank.GetAccount(id)
	if err := account.Withdraw(req.Amount); err != nil {
		a.ErrorInternal(w, err)
		return
	}

	a.Ok(w, nil)
}

// @summary Get account balance
// @tags account
// @accept json
// @produce json
// @param id path string true "a202482a-bf68-4e17-a4c4-5a268b2e15d6"
// @success 200 {object} responder.Response "Success"
// @failure 400 {object} responder.Response "Bad request"
// @router /accounts/{id}/balance [get]
func (a *Controller) GetBalance(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		a.ErrorBadRequest(w, fmt.Errorf("id is empty"))
		return
	}

	account := a.bank.GetAccount(id)
	balance := account.GetBalance()

	a.Ok(w, getBalanceResponse{Balance: balance})
}
