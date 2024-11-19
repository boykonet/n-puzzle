package priority_queue

import (
	"container/heap"
	"fmt"
)

type Item[T any] struct {
	value    T
	priority int
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue [T any] []*Item[T]

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leaks
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
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
	heap.Push(&q.pq, &Item[T]{value: value, priority: priority})
}

func (q *PriorityQueueImpl[T]) Dequeue() (T, error) {
	if q.Len() == 0 {
		var zero T
		return zero, fmt.Errorf("queue is empty")
	}
	item := heap.Pop(&q.pq).(*Item[T])
	return item.value, nil
}

func (q *PriorityQueueImpl[T]) Peek() (T, error) {
	if q.Len() == 0 {
		var zero T
		return zero, fmt.Errorf("queue is empty")
	}
	return q.pq[0].value, nil
}

func (q *PriorityQueueImpl[T]) Len() int {
	return q.pq.Len()
}
