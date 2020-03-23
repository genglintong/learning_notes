package main

import "sync"

var (
	//mu      sync.Mutex // guards balance
	mu	sync.RWMutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	b := balance
	return b
}
