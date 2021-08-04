// 3:Дана последовательность чисел  (2,4,6,8,10)
// найти их сумму квадратов(22+32+42….)
// с использованием конкурентных вычислений
package main

import "fmt"

func main() {
	// arr := [5]int{2,4,6,8,10}
	arr := []int{2, 4, 6, 8, 10}

	// Создание канала
	ch := make(chan int)

	// Ф-ция рассчитывает значения квадратов взятых из слайса arr
	go func(arr []int, ch chan int) {
		for _, v := range arr {
			ch <- v * v
		}
		close(ch)
	}(arr, ch)

	// Переменная хранения суммированных значений
	var sum int

	for v := range ch {
		//Сложить-присвоить в переменную
		sum += v
	}

	fmt.Println(sum)
}
