package queue

type Queue[T any] struct {
	elements []T
	size     int
	head     int
	tail     int
}

func New[T any](capacity int) *Queue[T] {
	return &Queue[T]{
		elements: make([]T, capacity),
	}
}

func (q *Queue[T]) Enqueue(item T) {
	if len(q.elements) == q.size {
		grow := make([]T, q.size*2)
		copy(grow, q.elements)
		q.elements = grow
	}
	q.elements[q.tail] = item
	q.size++
	q.tail++
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) Dequeue() T {
	actualHead := q.head
	q.head++
	q.size--
	return q.elements[actualHead]
}
