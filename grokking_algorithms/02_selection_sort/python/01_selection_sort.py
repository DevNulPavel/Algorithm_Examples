#! /usr/bin/env python2
# -*- coding: utf-8 -*-


# Ищем индекс наименьшего значение в массиве
def findSmallest(test_list):
    # Переменная с наименьшим значением
    smallest = test_list[0]
    # Индекс переменной с наименьшим значением
    smallest_index = 0

    # Перебираем все элементы и подбираем минимальное значение
    for i in range(1, len(test_list)):
        if test_list[i] < smallest:
            # Обновляем значения
            smallest_index = i
            smallest = test_list[i]

    # Возвращаем индекс
    return smallest_index

# Сортировка
def selectionSort(test_list):
    print(type(test_list))
    # Новый отсортированный список
    new_list = []
    for i in range(len(test_list)):
        # Находим индекс наименьшего элемента в массиве
        smallest_index = findSmallest(test_list)
        # извлекаем элемент по индексу
        smalest_value = test_list.pop(smallest_index)
        # Добавляем этот элемент в новый список
        new_list.append(smalest_value)

    return new_list

test_list = [5, 3, 6, 2, 10, 34, 99, 12, 135, 1233]
sorted_array = selectionSort(test_list)
print(sorted_array)
