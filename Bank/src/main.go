package main

import (
	"banking/app"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	newUser := app.NewStorage("Alice", 200)
	var wg sync.WaitGroup

	// Горутина для показа баланса
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			balance, err := newUser.GetBalance()
			if err != nil {
				fmt.Printf("[ОШИБКА БАЛАНСА] %v\n", err)
			}
			fmt.Printf("[БАЛАНС] Текущий счёт: %d руб.\n", balance)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Горутина для платежей
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			amount := rand.Intn(100) + 10
			err := newUser.Pay(amount)
			if err != nil {
				fmt.Printf("[ОШИБКА ПЛАТЕЖА] Не удалось снять %d руб.: %v\n", amount, err)
			} else {
				fmt.Printf("[ПЛАТЁЖ] Снято %d руб.\n", amount)
			}
			time.Sleep(700 * time.Millisecond)
		}
	}()

	// Горутина для пополнений
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			amount := rand.Intn(50) + 10
			err := newUser.Deposit(amount)
			if err != nil {
				fmt.Printf("[ОШИБКА ПОПОЛНЕНИЯ] Не удалось пополнить на %d руб.: %v\n", amount, err)
			} else {
				fmt.Printf("[ПОПОЛНЕНИЕ] Счёт пополнен на %d руб.\n", amount)
			}
			time.Sleep(600 * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Println("\n[ЗАВЕРШЕНО] Все операции завершены!")
}
