#! /usr/bin/env python2
# -*- coding: utf-8 -*-

def quicksort_recursive(array):
    if len(array) < 2:
        # Если массив 0го или 1го размера - сортировать не надо
        return array
    else:
        # Выбираем какое-то опорное значение
        pivot = array[0]
        # Формируем список элементов меньших, чем данный
        less = [i for i in array[1:] if i <= pivot]
        # Формируем список элементов больших, чем текущий
        greater = [i for i in array[1:] if i > pivot]

        # Выполняем рекурсивный вызов сортировок с
        # объединением первого списка, центрального элемента, и большего списка
        return quicksort(less) + [pivot] + quicksort(greater)


# This function is same in both iterative and recursive
def partition(arr,l,h): 
    i = ( l - 1 ) 
    x = arr[h] 
  
    for j in range(l , h): 
        if   arr[j] <= x: 
  
            # increment index of smaller element 
            i = i+1
            arr[i],arr[j] = arr[j],arr[i] 
  
    arr[i+1],arr[h] = arr[h],arr[i+1] 
    return (i+1) 

def quicksort_no_recursive(input_list): 
    arr = list(input_list)

    l = 0
    h = len(arr)-1

    # Создаем стек для промежуточных результатов
    size = h - l + 1
    stack = [0] * (size) 
  
    # Инициализируем вершину стека
    top = -1
  
    # push initial values of l and h to stack 
    top = top + 1
    stack[top] = l 
    top = top + 1
    stack[top] = h 
  
    # Keep popping from stack while is not empty 
    while top >= 0: 
  
        # Pop h and l 
        h = stack[top] 
        top = top - 1
        l = stack[top] 
        top = top - 1
  
        # Set pivot element at its correct position in 
        # sorted array 
        p = partition( arr, l, h ) 
  
        # If there are elements on left side of pivot, 
        # then push left side to stack 
        if p-1 > l: 
            top = top + 1
            stack[top] = l 
            top = top + 1
            stack[top] = p - 1
  
        # If there are elements on right side of pivot, 
        # then push right side to stack 
        if p+1 < h: 
            top = top + 1
            stack[top] = p + 1
            top = top + 1
            stack[top] = h 
    return arr


test_list = [10, 5, 2, 3]
new_list = quicksort_no_recursive(test_list)
print(new_list)

