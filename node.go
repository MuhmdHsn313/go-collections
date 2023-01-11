package collections

type Node[T comparable] struct {
	next *Node[T]
	prev *Node[T]
	Data T
}
