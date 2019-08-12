#! /usr/bin/env python2
# -*- coding: utf-8 -*-


# Пример работы жадного алгоритма

# Общий список необходимых штатов, которые надо покрыть
states_needed = set(["mt", "wa", "or", "id", "nv", "ut", "ca", "az"])

# Список станций, которые вещают в конкретных штатах.
stations = {}
stations["kone"] = set(["id", "nv", "ut"])
stations["ktwo"] = set(["wa", "id", "mt"])
stations["kthree"] = set(["or", "nv", "ca"])
stations["kfour"] = set(["nv", "ut"])
stations["kfive"] = set(["ca", "az"])

# Конечный набор станций, который будет удовлетворять нашим условиям
final_stations = set()

# Пока есть список станций, которые нам надо покрыть - идем по циклу
while states_needed:

	best_station = None
	states_covered = set()

	# Перебираем все станции и ищем те, которые покрывают наибольшее количество нужных нам станций
	for station, states in stations.items():
		# Пересечение нужных нам станций и текущего значения
		covered = states_needed & states

		# Если полученное количество станций больше, чем количество уже покрытых
		if len(covered) > len(states_covered):
			# Сохраняем данную станцию как наилучший вариант
			best_station = station
			states_covered = covered

	# Из списка нужных станций отнимаем найденную станцию
	states_needed -= states_covered

	# Добавляем станцию к списку найденных
	final_stations.add(best_station)

print(final_stations)
