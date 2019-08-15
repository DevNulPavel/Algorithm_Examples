// Implementation of a Hash Table with Separate Chaining
// with linked lists using Horner's hash function
// http://en.wikipedia.org/wiki/Hash_table#Separate_chaining_with_linked_lists
package ht

import (
	"errors"
	"github.com/arnauddri/algorithms/data-structures/linked-list"
	"math"
)

// Структура данных хэш-таблицы
type HashTable struct {
	Table    []*list.List
	Size     int
	Capacity int
}

// Внутренний итем таблицы
type item struct {
	key   string
	value interface{}
}

// Создание новой хэш-таблицы
func New(cap int) *HashTable {
	// Создаем внутренний массив
	internalArray := make([]*list.List, cap)
	// Создаем непосредственно сам объект
	newTable := HashTable{
		Table: internalArray, 
		Size: 0, 
		Capacity: cap,
	}
	return &newTable
}

// Метод получения итема из таблицы
func (ht *HashTable) Get(key string) (interface{}, error) {
	// Получаем индекс в массиве
	index := ht.position(key)
	// Получаем конкретный этэм из массива с нужным ключем
	item, err := ht.find(index, key)

	// Не нашли итем - возвращаем ошибку
	if item == nil {
		return "", errors.New("Not Found")
	}

	// Созвращаем итем и ошибку поиска, если есть
	return item.value, err
}

// Метод помещения объекта в таблицу
func (ht *HashTable) Put(key, value string) {
	// Вычисляем потенциальный индекс в таблице
	index := ht.position(key)

	// Проверяем, есть ли в массиве уже список по нужному нам индексу
	if ht.Table[index] == nil {
		// Если еще нету списка, то создаем его
		ht.Table[index] = list.NewList()
	}

	// Получаем итем под ключем, в котором хранится значение
	currentItem, err := ht.find(index, key)
	if err != nil {
		// Создаем внутренний итем
		newitem := &item{
			key: key, 
			value: value,
		}

		// Добавляем новый итем к таблице
		ht.Table[index].Append(newitem)
		ht.Size++
	} else {
		// Значение под данным ключем уже существует, поэтому просто обновляем значение
		currentItem.value = value
	}
}

func (ht *HashTable) Del(key string) error {
	// Получаем индекс в массиве
	index := ht.position(key)

	// Получаем список из таблицы, везвращаем ошибку если нету
	l := ht.Table[index]
	if l == nil {
		return errors.New("Not Found")
	}

	// Перебираем все итемы в списке и ищем с нужным ключем
	var val *item
	l.Each(func(node list.Node) {
		if node.Value.(*item).key == key {
			val = node.Value.(*item)
		}
	})

	// Нашли или нет?
	if val == nil {
		return nil
	}

	// Уменьшаем количество элеметов
	ht.Size--

	// Удаляем из списка элемент
	return l.Remove(val)
}

func (ht *HashTable) ForEach(f func(*item)) {
	for k := range ht.Table {
		if ht.Table[k] != nil {
			ht.Table[k].Each(func(node list.Node) {
				f(node.Value.(*item))
			})
		}
	}
}

// Вычисляем позицию в массиве по строке
func (ht *HashTable) position(s string) int {
	// Получаем индекс в массиве как хэш % размер_массива
	return hashCode(s) % ht.Capacity
}

// Получение элемента из списка
func (ht *HashTable) find(i int, key string) (*item, error) {
	// Список элементов по конкретному индексу элементы
	l := ht.Table[i]
	
	// Результат
	var val *item

	// Перебираем все элементы списка пока не найдем нужный нам элемент
	l.Each(func(node list.Node) {
		if node.Value.(*item).key == key {
			val = node.Value.(*item)
		}
	})

	// Если не нашли - выходим
	if val == nil {
		return nil, errors.New("Not Found")
	}

	return val, nil
}

// Метод хорнера для вычисления хэша от строки
func hashCode(s string) int {
	hash := int32(0)
	for i := 0; i < len(s); i++ {
		hash = int32(hash<<5-hash) + int32(s[i])
		hash &= hash
	}
	return int(math.Abs(float64(hash)))
}
