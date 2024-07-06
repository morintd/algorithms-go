package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDepthFirstSearch(t *testing.T) {
	t.Run("Should return path", func(t *testing.T) {
		graph := WeightedAdjacencyList{
			[]GraphEdge{{to: 1}},
			[]GraphEdge{{to: 2}},
			[]GraphEdge{{to: 0}},
		}

		path := DepthFirstSearch(graph, 0, 2)

		assert.Equal(t, path, []int{0, 1, 2})
	})
}
