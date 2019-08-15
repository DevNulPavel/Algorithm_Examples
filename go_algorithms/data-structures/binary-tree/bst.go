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

func NewNode(i int) *Node {
	return &Node{Value: i}
}

func (n *Node) Compare(m *Node) int {
	// Сравниваем с текущим, меньше или больше параметр нод по значению 
	if n.Value < m.Value {
		return -1
	} else if n.Value > m.Value {
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
			} else {
				// если у текущего года есть левый ребенок, 
				// то делаем новую итерацию и проверяем уже его
				curNode = curNode.Left
			}
		} else {
			// Если нет правого нода, то добавляем новый нод в качестве правого нода
			if curNode.Right == nil {
				curNode.Right = n
				n.Parent = curNode
				break
			} else {
				// Если правый нод есть, то продолжаем итерироваться дальше
				curNode = curNode.Right
			}
		}
	}
	t.Size++
}

// Ищем конкретный нод для значения
func (t *Tree) Search(key int) int {
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
			return curNode.Value
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

	h := t.Head
	n := &Node{Key: key}
	for h != nil {
		switch n.Compare(h) {
		case -1:
			parent = h
			h = h.Left
		case 1:
			parent = h
			h = h.Right
		case 0:
			if h.Left != nil {
				right := h.Right
				h.Value = h.Left.Value
				h.Left = h.Left.Left
				h.Right = h.Left.Right

				if right != nil {
					subTree := &Tree{Head: h}
					IterOnTree(right, func(n *Node) {
						subTree.Insert(n.Value)
					})
				}
				t.Size--
				return true
			}

			if h.Right != nil {
				h.Value = h.Right.Value
				h.Left = h.Right.Left
				h.Right = h.Right.Right

				t.Size--
				return true
			}

			if parent == nil {
				t.Head = nil
				t.Size--
				return true
			}

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
