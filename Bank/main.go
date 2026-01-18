package main

import (
	"banking/app"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func GorutuneDemo(newUser *app.Storage, userID int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("[ПОЛЬЗОВАТЕЛЬ %d] Начало работы\n", userID)
	startTime := time.Now()

	for {

		balance, err := newUser.GetBalance()
		if err != nil {
			fmt.Printf("[ПОЛЬЗОВАТЕЛЬ %d] Ошибка баланса: %v\n", userID, err)
		}
		fmt.Printf("[ПОЛЬЗОВАТЕЛЬ %d] Текущий баланс: %d\n", userID, balance)

		err = newUser.Pay(rand.Intn(100))
		if err != nil {
			fmt.Printf("[ПОЛЬЗОВАТЕЛЬ %d] Ошибка платежа: %v\n", userID, err)
		} else {
			fmt.Printf("[ПОЛЬЗОВАТЕЛЬ %d] Платёж выполнен\n", userID)
		}

		err = newUser.Deposit(rand.Intn(50))
		if err != nil {
			fmt.Printf("[ПОЛЬЗОВАТЕЛЬ %d] Ошибка пополнения: %v\n", userID, err)
		} else {
			fmt.Printf("[ПОЛЬЗОВАТЕЛЬ %d] Счёт пополнен\n", userID)
		}

		vas := rand.Intn(667)
		if vas == 67 || vas == 52 || vas == 42 || vas == 0 || vas == 666 || vas == 13 || vas == 7 || vas == 3 || vas == 1 || vas == 99 {
			fmt.Printf("[ПОЛЬЗОВАТЕЛЬ %d] 🎉 ВЫПАЛО ЧИСЛО: %d\n", userID, vas)

			break
		}
		fmt.Printf("[ПОЛЬЗОВАТЕЛЬ %d] Случайное число: %d\n", userID, vas)
	}

	elapsed := time.Since(startTime)
	fmt.Printf("[ПОЛЬЗОВАТЕЛЬ %d] Завершен! Время выполнения: %v\n\n", userID, elapsed)
}

func main() {

	newUser := app.NewStorage("Alice", 200)
	newUser2 := app.NewStorage("Bob", 2300)
	newUser3 := app.NewStorage("Charlie", 20450)
	newUser4 := app.NewStorage("Diana", 2100)

	var wg sync.WaitGroup
	startMain := time.Now()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("ЗАПУСК ПАРАЛЛЕЛЬНОГО ВЫПОЛНЕНИЯ 4 ГОРУТИН")
	fmt.Println("═══════════════════════════════════════════════════════════")

	wg.Add(1)
	go GorutuneDemo(newUser, 1, &wg)

	wg.Add(1)
	go GorutuneDemo(newUser2, 2, &wg)

	wg.Add(1)
	go GorutuneDemo(newUser3, 3, &wg)

	wg.Add(1)
	go GorutuneDemo(newUser4, 4, &wg)

	wg.Wait()

	totalTime := time.Since(startMain)
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Printf("ВСЕ ГОРУТИНЫ ЗАВЕРШЕНЫ! ОБЩЕЕ ВРЕМЯ ВЫПОЛНЕНИЯ: %v\n", totalTime)
	fmt.Println("═══════════════════════════════════════════════════════════")
}
