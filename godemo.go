package godemo

import (
	"cmp"
)

type OrderableFn[T any] func(t1, t2 T) int

type Tree[T any] struct {
	Root    *node[T]
	OrderFn OrderableFn[T]
}
type node[T any] struct {
	val   T
	left  *node[T]
	right *node[T]
}

func (t *Tree[T]) Insert(val T) {
	t.Root = t.Root.insert(val, t.OrderFn)
}

func (t *Tree[T]) Contains(val T) bool {
	return t.Root.contains(val, t.OrderFn)
}

func (n *node[T]) insert(val T, f OrderableFn[T]) *node[T] {
	if n == nil {
		return &node[T]{
			val: val,
		}
	}
	switch order := f(n.val, val); order {
	case -1:
		n.left = n.left.insert(val, f)
	case 1:
		n.right = n.right.insert(val, f)
	default:
		n.val = val
	}
	return n
}

func (n *node[T]) contains(val T, f OrderableFn[T]) bool {
	if n == nil {
		return false
	}
	switch order := f(n.val, val); order {
	case -1:
		return n.left.contains(val, f)
	case 1:
		return n.right.contains(val, f)
	default:
		return true
	}
}

func NewTree[T any](f OrderableFn[T]) *Tree[T] {
	return &Tree[T]{
		OrderFn: f,
	}
}

type Person struct {
	// structs w/ comparable underlying fields are also comparable (on value)
	Name string
	Age  int
}

func (person Person) Compare(other Person) int {
	order := cmp.Compare[string](person.Name, other.Name)
	if order == 0 {
		return cmp.Compare[int](person.Age, other.Age)
	}
	return order
}
