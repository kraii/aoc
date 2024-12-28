package aoc

import "container/heap"

type Queueable interface {
	Priority() int
}

type Item[T any] struct {
	contents T
	priority int
}

func (i Item[T]) Priority() int {
	return i.priority
}

type PriorityQueue[T Queueable] []T

func (pq *PriorityQueue[T]) Len() int { return len(*pq) }

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	q := *pq
	return q[i].Priority() < q[j].Priority()
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	q := *pq
	q[i], q[j] = q[j], q[i]
}

func (pq *PriorityQueue[T]) Push(x any) {
	*pq = append(*pq, x.(T))
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) Init() {
	heap.Init(pq)
}

func (pq *PriorityQueue[T]) PopHeap() T {
	return heap.Pop(pq).(T)
}

func (pq *PriorityQueue[T]) PushHeap(v T) {
	heap.Push(pq, v)
}

func PushItem[T any](pq *PriorityQueue[Item[T]], item T, priority int) {
	pq.PushHeap(Item[T]{contents: item, priority: priority})
}

func PopItem[T any](pq *PriorityQueue[Item[T]]) T {
	return pq.PopHeap().contents
}

func MakeQueue[T any](capacity int) *PriorityQueue[Item[T]] {
	pq := make(PriorityQueue[Item[T]], 0, capacity)
	pq.Init()
	return &pq
}
