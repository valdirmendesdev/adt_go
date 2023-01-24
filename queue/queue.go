package queue

type Queue[T any] struct {
	elements []T
	size     int
	tail     int
}

func New[T any](capacity int) *Queue[T] {
	return &Queue[T]{
		elements: make([]T, capacity),
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.elements[q.tail] = item
	q.size++
}

func (q *Queue[T]) Size() int {
	return q.size
}
