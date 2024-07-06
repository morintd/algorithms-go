package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBreadthFirstSearch(t *testing.T) {
	t.Run("Should return path", func(t *testing.T) {
		graph := WeightedAdjacencyMatrix{
			[]int{0, 1, 0},
			[]int{0, 0, 1},
			[]int{1, 0, 0},
		}

		path := BreadthFirstSearch(graph, 0, 2)

		assert.Equal(t, path, []int{0, 1, 2})
	})
}
