package main

import "fmt"

func main() {
	var amount float64
	fmt.Print("Введите сумму перевода: ")
	fmt.Scan(&amount)

	fmt.Printf("Вы ввели: %.2f сум\n", amount)
}
