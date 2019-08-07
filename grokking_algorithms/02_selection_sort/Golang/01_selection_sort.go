package main

import "fmt"

// Ищем индекс наименьшего значение в массиве
func findSmallest(arr []int) int {
	// Переменная с наименьшим значением
	smallest := arr[0]
	// Индекс переменной с наименьшим значением
	smallest_index := 0

	// Перебираем все элементы и подбираем минимальное значение
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}

	// Возвращаем индекс
	return smallest_index
}

// Сортировка
func selectionSort(arr []int) []int {
	// Размер входного массива
	size := len(arr)
	// Новый отсортированный список
	newArr := make([]int, size)
	for i := 0; i < size; i++ {
		// Находим индекс наименьшего элемента в массиве
		smallest := findSmallest(arr)
		// Сохраняем элемент в новом массиве
		newArr[i] = arr[smallest]
		// Тут вот шаткий момент, происходит создание нового массива из старого
		// можно оптимизировать этот момент?
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}
	return newArr
}

func main() {
	s := []int{5, 3, 6, 2, 10}
	fmt.Println(selectionSort(s))
}
