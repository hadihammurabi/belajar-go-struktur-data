package stack

import (
	"testing"

	"github.com/go-playground/assert"
)

var testData = []int{
	1,
	2,
	3,
	4,
	5,
}

func TestNewStack(t *testing.T) {
	stack := NewStack()

	assert.NotEqual(t, stack, nil)
	assert.Equal(t, stack.Size(), uint(0))
}

func TestAddData(t *testing.T) {
	stack := NewStack()

	for _, d := range testData {
		stack.Push(d)
	}

	assert.Equal(t, stack.Size(), uint(len(testData)))
}

func TestGetData(t *testing.T) {
	stack := NewStack()

	for _, d := range testData {
		stack.Push(d)
	}

	reversedTestData := make([]int, len(testData))
	copy(reversedTestData, testData)
	for i, j := 0, len(reversedTestData)-1; i < j; i, j = i+1, j-1 {
		reversedTestData[i], reversedTestData[j] = reversedTestData[j], reversedTestData[i]
	}

	for _, d := range reversedTestData {
		assert.Equal(t, stack.Pop(), d)
	}
}
