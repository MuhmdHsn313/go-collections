package main

import (
	"fmt"
	"github.com/MuhmdHsn313/go-collections"
)

type person struct {
	name string
	age  uint
}

func (p person) CompareTo(other person) int {
	if p.age > other.age {
		return 1
	} else if p.age < other.age {
		return -1
	}
	return 0
}

func main() {
	linkedlist := collections.NewLinkedList[person]()

	linkedlist.AddFirst(person{name: "Muhammad", age: 20})
	linkedlist.AddFirst(person{name: "Ahmed", age: 21})

	fmt.Println(linkedlist.Size())

	p := linkedlist.PeekFirst()
	fmt.Println(p.name, p.age)

	p = linkedlist.PeekLast()
	fmt.Println(p.name, p.age)

	p = linkedlist.RemoveFirst()
	fmt.Println(p.name, p.age)

	p = linkedlist.PeekFirst()
	fmt.Println(p.name, p.age)

	p = linkedlist.RemoveLast()
	fmt.Println(p.name, p.age)

	fmt.Println(linkedlist.Size())

	linkedlist.AddFirst(person{name: "Muhammad", age: 20})
	linkedlist.AddFirst(person{name: "Ahmed", age: 21})
	linkedlist.AddFirst(person{name: "Basim", age: 22})

	linkedlist.Reverse()

	linkedlist.ForEach(func(p person) {
		fmt.Println(p.name, p.age)
	})

	println(linkedlist.Contains(person{name: "Muhammad", age: 20}))
	println(linkedlist.Contains(person{name: "Muhammad", age: 21}))

	linkedlist.Clear()
}
