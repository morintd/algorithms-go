package basecasepathfinding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasePathfinding(t *testing.T) {
	t.Run("Should solve maze and return path", func(t *testing.T) {
		maze := [][]Square{
			{SquareWall, SquareWall, SquareWall, SquareWall, SquareWall},
			{SquareWall, SquareEmpty, SquareEmpty, SquareEmpty, SquareWall},
			{SquareWall, SquareWall, SquareWall, SquareEmpty, SquareWall},
			{SquareWall, SquareEnd, SquareEmpty, SquareEmpty, SquareWall},
			{SquareWall, SquareWall, SquareWall, SquareWall, SquareWall},
		}

		start := Point{1, 1}
		isEnd, path := Solve(maze, start)

		expected := []Point{{1, 1}, {2, 1}, {3, 1}, {3, 2}, {3, 3}, {2, 3}, {1, 3}}
		assert.Equal(t, isEnd, true)
		assert.Equal(t, path, expected)
	})

}
