package queue

type IQueue[T any] interface {
	Empty() bool
	Size() int
	Front() T
	Back() T
	Push(element T)
	Pop()
	//Swap()
}
