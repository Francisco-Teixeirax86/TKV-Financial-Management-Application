package main

import (
	"backend/data"
	"fmt"
	"net/http"
	"strconv"
)

var store *data.Store

func main() {

	//Initialize a new store
	store = data.NewStore()

	//TESTING - Create dummy accounts
	store.CreateAccount("acc1", 1000)
	store.CreateAccount("acc2", 160)

	http.HandleFunc("/balance", getBalanceHandler)
	http.HandleFunc("/deposit", depositHandler)
	http.HandleFunc("/withdraw", withdrawHandler)
	http.HandleFunc("/transfer", transferHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the TKV Financial App!")
	})

	fmt.Println("Server is running, listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}

// depositHandler Handler for the deposit request
func depositHandler(writer http.ResponseWriter, request *http.Request) {
	accountID := request.URL.Query().Get("accountID")
	amountStr := request.URL.Query().Get("amount")
	if amountStr == "" || accountID == "" {
		http.Error(writer, "AccountID and Amount are required to update a balance", http.StatusBadRequest)
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)

	if err != nil {
		http.Error(writer, "invalid amount", http.StatusInternalServerError)
		return
	}

	err = store.Deposit(accountID, amount)
	if err != nil {
		http.Error(writer, "failed to update balance", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(writer, "Account %s has been successfully updated with amoun %s", accountID, amount)

}

// getBalanceHandler Handler for the get balance request
func getBalanceHandler(writer http.ResponseWriter, request *http.Request) {
	accountID := request.URL.Query().Get("accountID")
	if accountID == "" {
		http.Error(writer, "Account ID is empty", http.StatusBadRequest)
		return
	}

	balance, err := store.GetBalance(accountID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintf(writer, "Account %s, Balance %.2f \n", accountID, balance)
}

func withdrawHandler(writer http.ResponseWriter, request *http.Request) {
	accountID := request.URL.Query().Get("accountID")
	amountStr := request.URL.Query().Get("amount")

	if amountStr == "" || accountID == "" {
		http.Error(writer, "AccountID and Amount are required to withdraw a balance", http.StatusBadRequest)
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(writer, "invalid amount", http.StatusInternalServerError)
	}

	err = store.Withdraw(accountID, amount)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(writer, "Withrew %.2f from Account: %s \n", amount, accountID)
}

func transferHandler(writer http.ResponseWriter, request *http.Request) {
	fromAccountID := request.URL.Query().Get("fromAccountID")
	toAccountID := request.URL.Query().Get("toAccountID")
	amountStr := request.URL.Query().Get("amount")

	if amountStr == "" || fromAccountID == "" || toAccountID == "" {
		http.Error(writer, "fromAccountID, toAccountID and amount are required to transfer", http.StatusBadRequest)
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(writer, "invalid amount", http.StatusBadRequest)
	}

	err = store.Transfer(fromAccountID, toAccountID, amount)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(writer, "Transfered %.2f from Account: %s to Account: %s\n", amount, fromAccountID, toAccountID)

}
