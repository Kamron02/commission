package main

import "fmt"

func main() {
	var amount float64
	fmt.Print("Введите сумму перевода: ")
	fmt.Scan(&amount)

	commission := amount * 0.0029 // 0.29% комиссия
	fmt.Printf("Комиссия: %.2f сум\n", commission)
}
