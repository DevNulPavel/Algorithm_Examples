#! /usr/bin/env python2
# -*- coding: utf-8 -*-

from collections import deque

# Функция проверки, что данный персонаж является продавцом
def person_is_seller(name):
    # Просто проверяем, что его имя заканчивается на m
    return name[-1] == 'm'

# Описание графа
graph = {}
graph["you"] = ["alice", "bob", "claire"]
graph["bob"] = ["anuj", "peggy"]
graph["alice"] = ["peggy"]
graph["claire"] = ["thom", "jonny"]
graph["anuj"] = []
graph["peggy"] = []
graph["thom"] = []
graph["jonny"] = []

# Конвертируем словарь родителей в путь от конечного нода до начального
def convert_parents_to_path(parents, start_node, end_node):
    result = []

    parent = end_node
    result.append(parent)

    complete = False
    while not complete:
        parent = parents[parent]
        result.append(parent)

        if parent == start_node:
            complete = True

    result.reverse()

    return result

# Алгоритм поиска в ширину
def search(name):
    # Очередь поиска
    search_queue = deque()
    # В самом начале добавляем соседние для стартового узлы в качестве нодов поиска
    childrens = graph[name]
    search_queue += childrens

    # Сохраняем обратный путь от дочернего нода к родителю
    path_parents = {}
    for child in childrens:
        path_parents[child] = name
    
    # Список уже посещенных узлов
    searched = set()

    # Повторяем, пока у нас не закончилась очередь поиска
    while search_queue:
        # Берем первый элемент для поиска
        person = search_queue.popleft()
        
        # Проверяем, что мы еще не использовали этот узел
        if person in searched:
            continue

        # Проверяем, не является ли данный нод подходящим
        if person_is_seller(person):
            # Возвращаем успешный результат и найденный путь
            return True, convert_parents_to_path(path_parents, name, person)

        else:
            # Если не дошли до конца, то добавляем в очередь поиска всех соседей текущего узла
            childrens = graph[person]
            search_queue += childrens
            
            # Сохраняем обратный путь от дочернего нода к родителю
            for child in childrens:
                path_parents[child] = person

            # Помечаем данный нод как уже обработанный
            searched.add(person)

    return False, None

result, path = search("you")
if result:
    print path
    # Если является подходящим, то выходим из нашей функции с результатом
    print(path[-1] + " is a mango seller!")

