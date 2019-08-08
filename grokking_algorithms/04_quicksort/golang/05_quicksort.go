package main

import "fmt"

func quicksort(list []int) []int {
	// Сортировать имеет только большие массивы
	if len(list) >= 2 {
		// Выбираем опорный элемент
		pivot := list[0]

		// Создаем списки для всех элементов больше и всех элементов меньше опорного
		var less = []int{}
		var greater = []int{}

		// Идем по списку, по всем элементам кроме опорного
		for _, num := range list[1:] {
			// Добавляем в нужные массивы
			if pivot > num {
				less = append(less, num)
			} else {
				greater = append(greater, num)
			}
		}

		// Рекурсивно сортируем элементы меньше текущего
		less = quicksort(less)
		// Добавляем опорный элемент
		less = append(less, pivot)
		// Сортируем элементы больше текущего
		greater = quicksort(greater)
		// Финальный массив
		return append(less, greater...)
	} else {
		return list
	}
}

func main() {
	fmt.Println(quicksort([]int{10, 5, 2, 3}))
}
