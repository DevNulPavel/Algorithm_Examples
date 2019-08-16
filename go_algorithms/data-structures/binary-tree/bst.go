package bst

import ()

// Узел дерева
type Node struct {
	Key  int   	    // Ключ
	Value  int   	// Переменная
	Parent *Node    // Родитель
	Left   *Node    // Левый ребенок
	Right  *Node    // Правый ребенок
}

func NewNode(key, val int) *Node {
	return &Node{Key: key, Value: val}
}

func (n *Node) Compare(m *Node) int {
	// Сравниваем с текущим, меньше или больше параметр нод по значению 
	if n.Key < m.Key {
		return -1
	} else if n.Key > m.Key {
		return 1
	} else {
		return 0
	}
}

/////////////////////////////////////////////////////////////////

// Описание дерева
type Tree struct {
	Head *Node     // Верхушка данного поддерева
	Size int       // Размер дерева
}

// Создаем новое дерево с узлом-верхушкой 
func NewTree(n *Node) *Tree {
	if n == nil {
		return &Tree{}
	}
	return &Tree{Head: n, Size: 1}
}

// Добавляем новый элемент в дерево
func (t *Tree) Insert(key, val int) {
	// Создаем новый узел
	n := &Node{Key: key, Value: val}

	// Если у текущего дерева нету еще верхушки, значит новый нод и есть верхушка
	// дальше нет смысла делать что-то
	if t.Head == nil {
		t.Head = n
		t.Size++
		return
	}

	// Иначе берем верхушку дерева
	curNode := t.Head

	// И ищем, куда бы впихнуть новый нод
	for {
		// Если новый нод меньше текущей верхушки
		if n.Compare(curNode) == -1 {
			// и если у текущего проверяемого нода нету левого чилда, значит новый нод будет тем самым чилдом
			if curNode.Left == nil {
				curNode.Left = n
				n.Parent = curNode
				break
			} else if curNode.Left.Key == key {
				// TODO: Здесь требуется доработка, если ключи совпадают - перезаписываем
				curNode.Left.Value = val;
			} else {
				// если у текущего года есть левый ребенок, 
				// то делаем новую итерацию и проверяем уже его
				curNode = curNode.Left
			}
		}  else {
			// Если нет правого нода, то добавляем новый нод в качестве правого нода
			if curNode.Right == nil {
				curNode.Right = n
				n.Parent = curNode
				break
			} else if curNode.Right.Key == key {
				// TODO: Здесь требуется доработка, если ключи совпадают - перезаписываем
				curNode.Right.Value = val;
			} else {
				// Если правый нод есть, то продолжаем итерироваться дальше
				curNode = curNode.Right
			}
		}
	}
	t.Size++
}

// Ищем конкретный нод для значения
func (t *Tree) Search(key int) *Node {
	// Берем верхушку дерева
	curNode := t.Head
	// Создаем временный нод
	tempNode := &Node{Key: key}

	for curNode != nil {
		// Сравниваем значение тестируемого нода, 
		// чтобы определить куда идти, затем обходим ноды
		switch curNode.Compare(tempNode) {
		case -1:
			curNode = curNode.Right
		case 1:
			curNode = curNode.Left
		case 0:
			return curNode
		default:
			panic("Node not found")
		}
	}
	panic("Node not found")
}

// returns true if a node with value i was found
// and deleted and returns false otherwise
func (t *Tree) Delete(key int) bool {
	var parent *Node

	// Начинаем поиск с вершины дерева
	curNode := t.Head

	// Создаем временный нод для поиска
	n := &Node{Key: key}

	// Выполняем поиск по дереву
	for curNode != nil {
		// Сравниваем тестовый нод с текущим и определяем в какую сторону надо идти
		switch n.Compare(curNode) {
		case -1:
			// Если тестовый нод меньше, идем налево,
			// выбираем в качестве родителя текущий нод, 
			// а в качестве нового тестируемого - текущую ветку
			parent = curNode
			curNode = curNode.Left
		case 1:
			// Если тестовый нод больше, идем налево,
			// выбираем в качестве родителя текущий нод, 
			// а в качестве нового тестируемого - текущую ветку
			parent = curNode
			curNode = curNode.Right
		case 0:
			// Если мы нашли такой же ключ

			// Если левый узел не равен нулю
			if curNode.Left != nil {
				// Сохраняем текущий правый узел
				right := curNode.Right

				// Обновляем значения текущего нода на те, которые у левого нода
				curNode.Key = curNode.Left.Key
				curNode.Value = curNode.Left.Value
				curNode.Left = curNode.Left.Left
				curNode.Right = curNode.Left.Right

				// Если у текущего нода было правое поддерево
				if right != nil {
					// То добавляем содержимое этого правого поднода к текущему обновленному ноду,
					// происходит переделка всего поддерева
					subTree := &Tree{Head: curNode}
					IterOnTree(right, func(n *Node) {
						subTree.Insert(n.Key, n.Value)
					})
				}
				t.Size--
				return true
			}

			// Если у текущего нода есть правое поддерево, при этом левого нету,
			// то просто обновляем содержимое текущего нода, на то, что было у правого
			if curNode.Right != nil {
				// Обнов
				curNode.Key = curNode.Right.Key
				curNode.Value = curNode.Right.Value
				curNode.Left = curNode.Right.Left
				curNode.Right = curNode.Right.Right

				t.Size--
				return true
			}

			// Если текущего нода нету - значит выходим
			if parent == nil {
				t.Head = nil
				t.Size--
				return true
			}

			// Если это это конечные узлы, то просто удаляем их
			if parent.Left == n {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
			t.Size--
			return true
		}
	}
	return false
}

// Функция добавления к ноду поддерева
func IterOnTree(n *Node, f func(*Node)) bool {
	if n == nil {
		return true
	}
	if !IterOnTree(n.Left, f) {
		return false
	}

	f(n)

	return IterOnTree(n.Right, f)
}
