package handlers

import (
	"PenzaTestTask/internal/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type accountHandler struct {
	service services.AccountService
}

type request struct {
	Amount float64 `json:"amount"`
}

func NewAccountHandler(service services.AccountService) AccountHandler {
	return &accountHandler{service: service}
}

func (h *accountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	account := h.service.CreateAccount()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func (h *accountHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resultChan := make(chan error)
	go func() {
		resultChan <- h.service.Deposit(id, req.Amount)
	}()

	if err := <-resultChan; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (h *accountHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resultChan := make(chan error)
	go func() {
		resultChan <- h.service.Withdraw(id, req.Amount)
	}()

	if err := <-resultChan; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (h *accountHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	resultChan := make(chan struct {
		balance float64
		err     error
	})
	go func() {
		balance, err := h.service.GetBalance(id)
		resultChan <- struct {
			balance float64
			err     error
		}{balance, err}
	}()

	result := <-resultChan
	if result.err != nil {
		http.Error(w, result.err.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]float64{"balance": result.balance})
	}
}
