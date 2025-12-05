package main

import "fmt"

func main() {

	var n int

	fmt.Print("VVedite kol-vo chifer: ")
	fmt.Scan(&n)

	// создаём пустой срез нужного размера
	nums := make([]int,n)

	// ввод значений
	fmt.Println("vvedi znacheniya:")
	for i :=0; i < n ; i++ {
		fmt.Scan(&nums[i])
	}

	// начальные значения для сумм и min/max
	sum :=0
	min := nums[0]
	max := nums[0]

	// summa



	// обработка всех чисел
	for _,v := range nums {
		sum += v

		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}
	fmt.Println("vveli kol-vo chifer: ", nums)
	fmt.Println("summa: ", sum)
	fmt.Println("min: ", min)
	fmt.Println("max", max)	

	/*nums := []int{1, 2, 3}
	fmt.Println("Начальный срез:", nums)
	fmt.Println("len:", len(nums), "cap:", cap(nums))

	nums = append(nums, 4)
	nums = append(nums, 5, 6, 7)

	fmt.Println("После append:", nums)
	fmt.Println("len:", len(nums), "cap:", cap(nums))

	fmt.Println("\nВывод элементов:")
	for i, v := range nums {
		fmt.Println(i, "→", v)
	}
		*/
}
