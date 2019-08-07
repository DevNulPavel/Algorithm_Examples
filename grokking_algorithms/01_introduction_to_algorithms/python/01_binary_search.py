#! /usr/bin/env python2
# -*- coding: utf-8 -*-


def binary_search(input_arr, item):
    # Начинаем с верхней и нижней границы входного массива
    low = 0
    high = len(input_arr) - 1

    # Пока диапазон не сузится до нуля - повторяем итерации
    while low <= high:
        # Получаем индекс элемента посередине
        mid = (low + high) // 2
        
        # Получаем значение по этому индексу
        guess = input_arr[mid]

        # Проверяем, тот ли это элемент, если тот - возвращаем индекс
        if guess == item:
            return mid

        if guess > item:
            # Значение больше нужного - проверяем большую половину
            high = mid - 1
        else:
            # Значение меньше нужного - проверяем нижнюю половину
            low = mid + 1

    # Такого итема не существует
    return None

# Создаем список отсортированных элементов
my_list = [1, 3, 5, 7, 9]

# Ищем индекс нужного элемента
index = binary_search(my_list, 3)
print(index) # => 1

# Ищем индекс нужного элемента
index = binary_search(my_list, 7)
print(index) # => 1

# Ищем несуществующий элемент, вернется None
index = binary_search(my_list, -1)
print(index) # => None
