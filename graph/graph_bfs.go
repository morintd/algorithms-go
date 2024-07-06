package graph

import "sort"

type WeightedAdjacencyMatrix [][]int

func BreadthFirstSearch(graph WeightedAdjacencyMatrix, source int, target int) []int {
	seen := make([]bool, len(graph))
	prev := make([]int, len(graph))

	for i := range prev {
		prev[i] = -1
	}

	seen[source] = true
	queue := []int{source}

	for len(queue) > 0 {
		current := queue[0]

		if len(queue) > 1 {
			queue = queue[1 : len(queue)-1]
		} else {
			queue = []int{}
		}

		if current == target {
			break
		}

		adjencencies := graph[current]
		for i := 0; i < len(adjencencies); i++ {
			if adjencencies[i] == 0 {
				continue
			}

			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = current

			queue = append(queue, i)
		}

		seen[current] = true
	}

	if prev[target] == -1 {
		return []int{}
	}

	current := target
	out := make([]int, 0)

	for prev[current] != -1 {
		out = append(out, current)
		current = prev[current]
	}

	sort.SliceStable(out, func(i, j int) bool {
		return true
	})

	return append([]int{source}, out...)
}
