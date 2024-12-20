package main

import (
	"fmt"
	"math/rand"
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
	defer wg.Done()

	for i := 0; i < 5; i++ {
		transaction := Transaction{
			ID:        sourceID*1000 + 1,
			Amount:    rand.Float64() * 100,
			Currency:  "USD",
			Timestamp: time.Now(),
		}

		fmt.Printf("Fetched Transaction %d from source %d\n", transaction.ID, sourceID)
		transactions <- transaction
		time.Sleep(time.Millisecond * 500)
	}
}

func TransformData(transaction Transaction) Transaction {
	if transaction.Currency == "USD" {
		transaction.Amount *= 0.9
		transaction.Currency = "EUR"
		fmt.Printf("Transformed transaction %d to %.2f %s \n", transaction.ID, transaction.Amount, transaction.Currency)
	}
}

func Worker(id int, transactions <-chan Transaction, store *DataStore, wg *sync.WaitGroup) {
	defer wg.Done()
	for transaction := range transactions {
		transformedTransaction := TransformData(transaction)
		store.AddTransaction(transformedTransaction)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	transactions := make(chan Transaction, 10)
	store := &DataStore{}
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go FetchData(i, transactions, &wg)
	}
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go Worker(i, transactions, store, &wg)
	}
	go func() {
		wg.Wait()
		close(transactions)
	}()
	time.Sleep(time.Second * 5)
	fmt.Println("Data Ingestion and transformation pipeline completed.")
}
