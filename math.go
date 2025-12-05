package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Println("Сумма:", a+b)
	fmt.Println("Разница:", a-b)
	fmt.Println("Произведение:", a*b)

	// правильное деление с float
	result := float64(a) / float64(b)
	fmt.Println("Деление:", result)
}
