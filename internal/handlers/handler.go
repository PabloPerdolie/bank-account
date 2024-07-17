package handlers

import "net/http"

type AccountHandler interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
	Deposit(w http.ResponseWriter, r *http.Request)
	Withdraw(w http.ResponseWriter, r *http.Request)
	GetBalance(w http.ResponseWriter, r *http.Request)
}
