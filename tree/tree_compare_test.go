package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeCompare(t *testing.T) {
	t.Run("Compare", func(t *testing.T) {
		t.Run("Should return true if trees equivalent", func(t *testing.T) {
			first := BinaryNode[int]{Value: 1,
				Left: &BinaryNode[int]{Value: 5,
					Left: &BinaryNode[int]{Value: 6},
				},
				Right: &BinaryNode[int]{Value: 10,
					Left: &BinaryNode[int]{Value: 11},
				},
			}

			second := BinaryNode[int]{Value: 1,
				Left: &BinaryNode[int]{Value: 5,
					Left: &BinaryNode[int]{Value: 6},
				},
				Right: &BinaryNode[int]{Value: 10,
					Left: &BinaryNode[int]{Value: 11},
				},
			}

			equivalent := Compare(&first, &second)
			assert.Equal(t, true, equivalent)
		})

		t.Run("Should return false if trees are not equivalent", func(t *testing.T) {
			first := BinaryNode[int]{Value: 1,
				Left: &BinaryNode[int]{Value: 5,
					Left: &BinaryNode[int]{Value: 6},
				},
				Right: &BinaryNode[int]{Value: 10,
					Left: &BinaryNode[int]{Value: 11},
				},
			}

			second := BinaryNode[int]{Value: 1,
				Left: &BinaryNode[int]{Value: 5,
					Left: &BinaryNode[int]{Value: 6},
				},
				Right: &BinaryNode[int]{Value: 10,
					Left:  &BinaryNode[int]{Value: 11},
					Right: &BinaryNode[int]{Value: 12},
				},
			}

			equivalent := Compare(&first, &second)
			assert.Equal(t, false, equivalent)
		})
	})
}
