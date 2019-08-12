#! /usr/bin/env python2
# -*- coding: utf-8 -*-


infinity = float("inf")

# Описание графа
graph = {}
graph["start"] = {"a": 6, "b": 2 }
graph["a"] = {"fin": 1}
graph["b"] = {}
graph["b"] = {"a": 3, "fin": 5}
graph["fin"] = {}

# Таблица стоимостей перехода к конкретному узлу
costs = {}
costs["a"] = 6
costs["b"] = 2
costs["fin"] = infinity

# Таблица переходов от чилда к родителю, нужна для восстановления обратного пути от конца в начало
parents = {}
parents["a"] = "start"
parents["b"] = "start"
parents["fin"] = None

# Список уже обработанных нодов
processed = set()

# Конвертируем словарь родителей в путь от конечного нода до начального
def convert_parents_to_path(parents, start_node, end_node):
    result = []

    total_weight = 0

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

# Функция поиска ближайшего по стоимости нода с учетом уже обработанных
def find_lowest_cost_node(costs):
    # Значение наименьшей стоимости
    lowest_cost = float("inf")
    lowest_cost_node = None

    # Обходим все ноды и ищем нод с наименьшей стоимостью и не обработанный
    for node in costs:
        cost = costs[node]
        # Если нод с наименьшей стоиместью, чем последний + не был обработан
        if (cost < lowest_cost) and (node not in processed):
            # Обновляем значения
            lowest_cost = cost
            lowest_cost_node = node
    # Возвращаем полученное значение
    return lowest_cost_node

def find_lowest_path():
    # Ищем нод с наименьшей стоимостью из необработанных
    node = find_lowest_cost_node(costs)

    # Если такой нод нашелся - начинаем поиск
    while node is not None:
        # Получаем текущую стоимость перехода к данному ноду (включая косвенный переход)
        cost = costs[node]

        # Получаем список соседей данного нода
        neighbors = graph[node]

        # Перебираем список соседей текущего нода
        for n in neighbors.keys():
            # Вычисляем новую суммарную стоимость перехода к текущему ноду
            new_cost = cost + neighbors[n]
            # Если предыдущая стоимость перехода к текущему ноду выше, чем текущая новая,
            # то имеет смысл обновить новую стоимость перехода
            if costs[n] > new_cost:
                costs[n] = new_cost
                # так как переход дешевле, то обновляем родителя для обратного разворачивания
                parents[n] = node
        # Помечаем нод как обработанный
        processed.add(node)
        # Заново ищем нод с наименьшей стоимостью, который мы еще не обработали
        node = find_lowest_cost_node(costs)

    print("Cost from the start to each node:")
    print(costs)
    path = convert_parents_to_path(parents, "start", "fin")
    print path
    print("Total path weight: " + str(costs["fin"]))


find_lowest_path();

