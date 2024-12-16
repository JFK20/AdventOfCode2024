package mathUtil

import (
	"math"
	"sort"
)

type PriorityQueue struct {
	elements []queueElement
}

type queueElement struct {
	pos      Vector2D[int]
	priority int
}

func (pq *PriorityQueue) Push(pos Vector2D[int], priority int) {
	pq.elements = append(pq.elements, queueElement{pos, priority})
	// Simple bubble up to maintain priority
	sort.Slice(pq.elements, func(i, j int) bool {
		return pq.elements[i].priority < pq.elements[j].priority
	})
}

func (pq *PriorityQueue) Pop() (Vector2D[int], int) {
	if len(pq.elements) == 0 {
		return Vector2D[int]{}, math.MaxInt32
	}

	item := pq.elements[0]
	pq.elements = pq.elements[1:]
	return item.pos, item.priority
}

func (pq *PriorityQueue) Len() int {
	return len(pq.elements)
}
