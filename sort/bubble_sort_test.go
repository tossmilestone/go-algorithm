package sort_test


import (
	"testing"
	sort "github.com/tossmilestone/go-algorithm/sort"
	assert "github.com/stretchr/testify/assert"
)


func TestBubbleSort(t *testing.T) {
	data := []int{3, 2, 1, 5, 6}
	expected := []int{1, 2, 3, 5, 6}
	sort.BubbleSort(data)
	for i := range data {
		assert.Equal(t, expected[i], data[i])
	}
}
