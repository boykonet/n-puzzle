package priority_queue

import (
	"container/heap"
	"errors"
)

var (
	ErrorEmptyQueue = errors.New("queue is empty")
)

type item[T any] struct {
	Value    T
	Priority int
	Index    int // The Index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue[T any] []*item[T]

func (pq *PriorityQueue[T]) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, Priority so we use greater than here.
	return (*pq)[i].Priority < (*pq)[j].Priority
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].Index = i
	(*pq)[j].Index = j
}

func (pq *PriorityQueue[T]) Push(x interface{}) {
	n := len(*pq)
	elem := x.(*item[T])
	elem.Index = n
	*pq = append(*pq, elem)
}

func (pq *PriorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	elem := old[n-1]
	old[n-1] = nil  // avoid memory leaks
	elem.Index = -1 // for safety
	*pq = old[0 : n-1]
	return elem
}

type PriorityQueueImpl[T any] struct {
	pq PriorityQueue[T]
}

func NewPriorityQueue[T any]() IQueue[T] {
	pq := make(PriorityQueue[T], 0)
	heap.Init(&pq)
	return &PriorityQueueImpl[T]{pq: pq}
}

func (q *PriorityQueueImpl[T]) Enqueue(value T, priority int) {
	heap.Push(&q.pq, &item[T]{Value: value, Priority: priority})
}

func (q *PriorityQueueImpl[T]) Dequeue() (T, error) {
	if q.Len() == 0 {
		var zero T
		return zero, ErrorEmptyQueue
	}
	elem := heap.Pop(&q.pq).(*item[T])
	return elem.Value, nil
}

func (q *PriorityQueueImpl[T]) Peek() (T, error) {
	if q.Len() == 0 {
		var zero T
		return zero, ErrorEmptyQueue
	}
	return q.pq[0].Value, nil
}

func (q *PriorityQueueImpl[T]) Len() int {
	return q.pq.Len()
}
