package queue

type queueStruct[T any] struct {
	array []T
}

func NewQueue[T any]() IQueue[T] {
	return &queueStruct[T]{
		array: make([]T, 0),
	}
}

func (q *queueStruct[T]) Empty() bool {
	return len(q.array) == 0
}

func (q *queueStruct[T]) Size() int {
	return len(q.array)
}

func (q *queueStruct[T]) Front() T {
	return q.array[0]
}

func (q *queueStruct[T]) Back() T {
	return q.array[len(q.array)-1]
}

func (q *queueStruct[T]) Push(element T) {
	q.array = append(q.array, element)
}

func (q *queueStruct[T]) Pop() {
	q.array = q.array[:len(q.array)-1]
}
