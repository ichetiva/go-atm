package controller

type createAccountResponse struct {
	ID string `json:"id"`
}

type operationRequest struct {
	Amount float64 `json:"amount"`
}

type getBalanceResponse struct {
	Balance float64 `json:"amount"`
}
