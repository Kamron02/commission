package main

import "fmt"

func main() {
	var amount float64
	fmt.Print("Введите сумму перевода: ")
	fmt.Scan(&amount)

	if amount < 500 {
		fmt.Println("Ошибка: Минимальная сумма перевода — 500 сум.")
		return
	}
	if amount > 15000000 {
		fmt.Println("Ошибка: Максимальная сумма перевода — 15 000 000 сум.")
		return
	}

	commission := amount * 0.0029 // 0.29% комиссия
	fmt.Printf("Комиссия: %.2f сум\n", commission)
}
