package queue

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

var testData = []int{
	1,
	2,
	3,
	4,
	5,
}

func TestNewQueue(t *testing.T) {
	queue := NewQueue()
	assert.NotEqual(t, queue, nil)
	assert.Equal(t, queue.Size(), uint(0))
}

func TestAddData(t *testing.T) {
	queue := NewQueue()
	for _, data := range testData {
		queue.Enqueue(data)
	}

	assert.Equal(t, queue.Size(), uint(len(testData)))
}

func TestGetData(t *testing.T) {
	queue := NewQueue()
	for _, data := range testData {
		queue.Enqueue(data)
	}

	for _, data := range testData {
		assert.Equal(t, queue.Dequeue(), data)
	}
}
