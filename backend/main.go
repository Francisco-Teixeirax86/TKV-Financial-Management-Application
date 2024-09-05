package main

import (
	"backend/data"
	"fmt"
	"net/http"
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

func depositHandler(writer http.ResponseWriter, request *http.Request) {
	accountID := request.URL.Query().Get("accountID")
	amount := request.URL.Query().Get("amount")
	if amount == "" || accountID == "" {
		http.Error(writer, "AccountID and Amount are required to update a balance", http.StatusBadRequest)
		return
	}

	depositAmount := 100.0 //FOR TESTING  ONLY
	err := store.UpdateBalance(accountID, depositAmount)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(writer, "Account %s has been successfully updated with amoun %s", accountID, depositAmount)

}

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

	fmt.Fprintf(writer, "Account %s, Balance %.2f\n", accountID, balance)
}
