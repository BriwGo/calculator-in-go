package app

import (
	"math/rand"
	"sync"
)

type Storage struct {
	mu      sync.Mutex // Мьютекс для защиты доступа
	ID      int
	Name    string
	Balance int
}

func NewStorage(name string, balance int) *Storage {
	return &Storage{
		ID:      rand.Int(),
		Name:    name,
		Balance: balance,
	}
}

func (l *Storage) GetBalance() (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.Balance, nil
}

func (l *Storage) Pay(amount int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	randomValue := rand.Intn(100)

	if randomValue < 30 {
		return NetworkError
	}

	if amount > l.Balance {
		return insufficientFunds
	}

	l.Balance -= amount
	return nil
}

func (l *Storage) Deposit(amount int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if amount < 0 {
		return WhatTheFuck
	}
	l.Balance += amount
	return nil
}
