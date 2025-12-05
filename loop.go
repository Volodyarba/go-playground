package main

import "fmt"

func main() {
	var n int

    fmt.Println("n = ")
	fmt.Scanln(&n)

	fmt.Println("\nvivod ot 1 do n:")
    for i := 1; i <= n; i++ {
        fmt.Println(i)
    }
	
	sum:=0
    for i := 1; i <= n; i++ {
		sum +=i
    }
	fmt.Println("summ = ",sum)

	fact :=1
    for i := 1; i <= n; i++ {
        fact *=i
    }
	fmt.Println("factorial = ",fact)
}
