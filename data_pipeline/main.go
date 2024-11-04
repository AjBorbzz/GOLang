package main

import (
	"fmt"
	"sync"
	"time"
)

type Transaction struct {
	ID        int
	Amount    float64
	Currency  string
	Timestamp time.Time
}

type DataStore struct {
	mu           sync.Mutex
	Transactions []Transaction
}

func (ds *DataStore) AddTransaction(t Transaction) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.Transactions = append(ds.Transactions, t)
	fmt.Printf("Stored transaction %d with amount %.2f %s\n", t.ID, t.Amount, t.Currency)
}

func FetchData(sourceID int, transactions chan<- Transaction, wg *sync.WaitGroup) {

}
