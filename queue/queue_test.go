package queue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdirmendesdev/adt_go/queue"
)

func TestQueue_Enqueue(t *testing.T) {
	q := queue.New[int](3)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Equal(t, 3, q.Size())
}

func TestQueue_EnqueueOverCapacity(t *testing.T) {
	q := queue.New[int](1)

	q.Enqueue(1)
	assert.Equal(t, 1, q.Size())
	q.Enqueue(2)
	assert.Equal(t, 2, q.Size())

	v, _ := q.Dequeue()
	assert.Equal(t, 1, v)
	v, _ = q.Dequeue()
	assert.Equal(t, 2, v)
}

func TestQueue_Dequeue(t *testing.T) {
	q := queue.New[string](3)
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")

	assert.Equal(t, 3, q.Size())
	value, _ := q.Dequeue()
	assert.Equal(t, "a", value)
	assert.Equal(t, 2, q.Size())

	value, _ = q.Dequeue()
	value, _ = q.Dequeue()
	assert.Equal(t, "c", value)
	assert.Equal(t, 0, q.Size())
}

func TestQueue_DequeueEmptyQueue(t *testing.T) {
	q := queue.New[int](1)
	_, err := q.Dequeue()
	assert.Error(t, err, "queue is empty")
}
