package main

import (
	"backend/data"
	"fmt"
	"sync"
)

func main() {
	store := data.NewStore()
	store.CreateAccount("account1", 1000)

	var wg sync.WaitGroup
	numGoroutines := 100
	depositAmount := 10.0

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := store.Deposit("account1", depositAmount)
			if err != nil {
				fmt.Print("Error: %s \n", err)
			}
		}()
	}

	wg.Wait()

	balance, _ := store.GetBalance("account1")
	expectedBalance := 1000 + float64(numGoroutines)*depositAmount
	fmt.Printf("Final balance: %.2f (Expected %.2f)\n", balance, expectedBalance)
}
