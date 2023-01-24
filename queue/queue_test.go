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
