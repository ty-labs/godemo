package godemo

import (
	"cmp"
	"fmt"
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

/*
Chapter exercises:

1) A generic function that will double any int/float
2) A generic interface that requies underlying int/float64 & embeds fmt.Stringer
+ a concrete type that meets the interfaces
3) A generic singly-linked LL data structure holding comparable values
*/
type Numeric interface {
	int | int8 | uint8 | int16 | uint16 | int32 | uint32 |
		int64 | uint64 | float32 | float64
}

func Double[T Numeric](n T) T {
	return 2 * n
}

type PrintableIntFloat interface {
	~int | ~float64
	fmt.Stringer
}

type PrintableFloat float64

func (pf PrintableFloat) String() string {
	return fmt.Sprintf("%.2f", pf)
}

func PrintPrintable[T PrintableIntFloat](pif T) {
	fmt.Println(pif)
}

type SinglyLinkedList[T comparable] struct {
	Len  int
	Root *singlyLinkedListNode[T]
}
type singlyLinkedListNode[T comparable] struct {
	val  T
	next *singlyLinkedListNode[T]
}

func (sll *SinglyLinkedList[T]) Add(val T) {
	// recursive wrappers... this could be done iteratively as well
	sll.Root = sll.Root.add(val)
	sll.Len += 1
}

func (sll *SinglyLinkedList[T]) Insert(val T, i int) {
	sll.Root = sll.Root.insert(val, i, 0)
	sll.Len += 1
}

func (sll *SinglyLinkedList[T]) Index(val T) int {
	return sll.Root.index(val, 0)
}

func (n *singlyLinkedListNode[T]) add(val T) *singlyLinkedListNode[T] {
	if n == nil {
		return &singlyLinkedListNode[T]{val: val}
	}
	n.next = n.next.add(val)
	return n
}

func (n *singlyLinkedListNode[T]) insert(val T, i int, curr int) *singlyLinkedListNode[T] {
	if n == nil {
		return &singlyLinkedListNode[T]{val: val}
	}
	if i == curr {
		newNode := singlyLinkedListNode[T]{val: val}
		newNode.next = n
		return &newNode
	}
	n.next = n.next.insert(val, i, curr+1)
	return n
}

func (n *singlyLinkedListNode[T]) index(val T, curr int) int {
	if n == nil {
		return -1
	}
	if n.val == val {
		return curr
	}
	return n.next.index(val, curr+1)
}

func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		Len: 0,
	}
}
