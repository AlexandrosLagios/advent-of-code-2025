package utils

import (
	"container/heap"
	"math"
)

// BFS performs breadth-first search
func BFS[T comparable](start T, neighbors func(T) []T, isGoal func(T) bool) ([]T, bool) {
	queue := NewQueue[T]()
	queue.Enqueue(start)
	visited := NewSet[T]()
	visited.Add(start)
	parent := make(map[T]T)

	for !queue.IsEmpty() {
		current, _ := queue.Dequeue()

		if isGoal(current) {
			// Reconstruct path
			path := []T{current}
			for {
				p, ok := parent[current]
				if !ok {
					break
				}
				path = append([]T{p}, path...)
				current = p
			}
			return path, true
		}

		for _, next := range neighbors(current) {
			if !visited.Contains(next) {
				visited.Add(next)
				parent[next] = current
				queue.Enqueue(next)
			}
		}
	}

	return nil, false
}

// DFS performs depth-first search
func DFS[T comparable](start T, neighbors func(T) []T, isGoal func(T) bool) ([]T, bool) {
	stack := NewStack[T]()
	stack.Push(start)
	visited := NewSet[T]()
	visited.Add(start)
	parent := make(map[T]T)

	for !stack.IsEmpty() {
		current, _ := stack.Pop()

		if isGoal(current) {
			// Reconstruct path
			path := []T{current}
			for {
				p, ok := parent[current]
				if !ok {
					break
				}
				path = append([]T{p}, path...)
				current = p
			}
			return path, true
		}

		for _, next := range neighbors(current) {
			if !visited.Contains(next) {
				visited.Add(next)
				parent[next] = current
				stack.Push(next)
			}
		}
	}

	return nil, false
}

// PriorityItem represents an item with a priority
type PriorityItem[T any] struct {
	Value    T
	Priority int
	Index    int
}

// PriorityQueue implements a min-heap priority queue
type PriorityQueue[T any] []*PriorityItem[T]

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue[T]) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PriorityItem[T])
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

// Dijkstra performs Dijkstra's shortest path algorithm
func Dijkstra[T comparable](start T, neighbors func(T) []T, cost func(T, T) int, isGoal func(T) bool) ([]T, int, bool) {
	pq := make(PriorityQueue[T], 0)
	heap.Init(&pq)

	heap.Push(&pq, &PriorityItem[T]{Value: start, Priority: 0})

	distances := map[T]int{start: 0}
	parent := make(map[T]T)
	visited := NewSet[T]()

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*PriorityItem[T])
		current := item.Value

		if visited.Contains(current) {
			continue
		}
		visited.Add(current)

		if isGoal(current) {
			// Reconstruct path
			path := []T{current}
			for {
				p, ok := parent[current]
				if !ok {
					break
				}
				path = append([]T{p}, path...)
				current = p
			}
			return path, distances[item.Value], true
		}

		for _, next := range neighbors(current) {
			if visited.Contains(next) {
				continue
			}

			newDist := distances[current] + cost(current, next)
			oldDist, exists := distances[next]

			if !exists || newDist < oldDist {
				distances[next] = newDist
				parent[next] = current
				heap.Push(&pq, &PriorityItem[T]{Value: next, Priority: newDist})
			}
		}
	}

	return nil, math.MaxInt, false
}

// AStar performs A* search
func AStar[T comparable](start T, neighbors func(T) []T, cost func(T, T) int, heuristic func(T) int, isGoal func(T) bool) ([]T, int, bool) {
	pq := make(PriorityQueue[T], 0)
	heap.Init(&pq)

	heap.Push(&pq, &PriorityItem[T]{Value: start, Priority: heuristic(start)})

	gScore := map[T]int{start: 0}
	parent := make(map[T]T)
	visited := NewSet[T]()

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*PriorityItem[T])
		current := item.Value

		if visited.Contains(current) {
			continue
		}
		visited.Add(current)

		if isGoal(current) {
			// Reconstruct path
			path := []T{current}
			for {
				p, ok := parent[current]
				if !ok {
					break
				}
				path = append([]T{p}, path...)
				current = p
			}
			return path, gScore[item.Value], true
		}

		for _, next := range neighbors(current) {
			if visited.Contains(next) {
				continue
			}

			tentativeG := gScore[current] + cost(current, next)
			oldG, exists := gScore[next]

			if !exists || tentativeG < oldG {
				gScore[next] = tentativeG
				parent[next] = current
				fScore := tentativeG + heuristic(next)
				heap.Push(&pq, &PriorityItem[T]{Value: next, Priority: fScore})
			}
		}
	}

	return nil, math.MaxInt, false
}
