package handlers

import "github.com/gorilla/mux"

func SetupRoutes(handler AccountHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/accounts", handler.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{id}/deposit", handler.Deposit).Methods("POST")
	r.HandleFunc("/accounts/{id}/withdraw", handler.Withdraw).Methods("POST")
	r.HandleFunc("/accounts/{id}/balance", handler.GetBalance).Methods("GET")

	return r
}
