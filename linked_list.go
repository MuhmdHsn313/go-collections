package collections

import (
	"bytes"
	"fmt"
	"sync"
)

type LinkedList[T comparable] struct {
	head  *Node[T]
	tail  *Node[T]
	size  int
	mutex sync.RWMutex
}

func NewLinkedList[T comparable]() LinkedList[T] {
	return LinkedList[T]{
		head:  nil,
		tail:  nil,
		size:  0,
		mutex: sync.RWMutex{},
	}
}

func (l *LinkedList[T]) Size() int {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	return l.size
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.Size() == 0
}

func (l *LinkedList[T]) AddFirst(data T) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	node := &Node[T]{
		next: l.head,
		prev: nil,
		Data: data,
	}

	if l.head != nil {
		l.head.prev = node
	}

	l.head = node

	if l.tail == nil {
		l.tail = node
	}

	l.size++
}

func (l *LinkedList[T]) AddLast(data T) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	node := &Node[T]{
		next: nil,
		prev: l.tail,
		Data: data,
	}

	if l.tail != nil {
		l.tail.next = node
	}

	l.tail = node

	if l.head == nil {
		l.head = node
	}

	l.size++
}

func (l *LinkedList[T]) RemoveFirst() *T {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.head == nil {
		return nil
	}

	data := l.head.Data
	l.head = l.head.next

	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}

	l.size--

	return &data
}

func (l *LinkedList[T]) RemoveLast() *T {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.tail == nil {
		return nil
	}

	data := l.tail.Data
	l.tail = l.tail.prev

	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}

	l.size--

	return &data
}

func (l *LinkedList[T]) PeekFirst() *T {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	if l.head == nil {
		return nil
	}

	data := l.head.Data
	return &data
}

func (l *LinkedList[T]) PeekLast() *T {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	if l.tail == nil {
		return nil
	}

	data := l.tail.Data
	return &data
}

func (l *LinkedList[T]) Get(index int) *T {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	if index >= l.size {
		return nil
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	data := current.Data
	return &data
}

func (l *LinkedList[T]) Remove(index int) *T {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if index >= l.size {
		return nil
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	if current.prev == nil {
		l.head = current.next
	} else {
		current.prev.next = current.next
	}

	if current.next == nil {
		l.tail = current.prev
	} else {
		current.next.prev = current.prev
	}

	l.size--

	data := current.Data
	return &data
}

func (l *LinkedList[T]) Clear() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *LinkedList[T]) Values() []T {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	values := make([]T, l.size)
	current := l.head
	for i := 0; current != nil; i++ {
		values[i] = current.Data
		current = current.next
	}

	return values
}

func (l *LinkedList[T]) String() string {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	var buffer bytes.Buffer
	current := l.head
	for current != nil {
		buffer.WriteString(fmt.Sprintf("%v", current.Data))
		current = current.next
		if current != nil {
			buffer.WriteString(" -> ")
		}
	}

	return buffer.String()
}

func (l *LinkedList[T]) Reverse() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	current := l.head
	for current != nil {
		current.next, current.prev = current.prev, current.next
		current = current.prev
	}

	l.head, l.tail = l.tail, l.head
}

func (l *LinkedList[T]) ForEach(callback func(data T)) bool {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	current := l.head
	for current != nil {
		callback(current.Data)
		current = current.next
	}

	return true
}

func (l *LinkedList[T]) ForEachReverse(callback func(data T)) bool {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	current := l.tail
	for current != nil {
		callback(current.Data)
		current = current.prev
	}

	return true
}

func (l *LinkedList[T]) Contains(data T) bool {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	current := l.head
	for current != nil {
		if current.Data == data {
			return true
		}
		current = current.next
	}

	return false
}

func (l *LinkedList[T]) IndexOf(data T) int {
	l.mutex.RLock()
	defer l.mutex.RUnlock()

	current := l.head
	for i := 0; current != nil; i++ {
		if current.Data == data {
			return i
		}
		current = current.next
	}

	return -1
}
