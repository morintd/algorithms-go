package graph

type GraphEdge struct {
	to int
}

type WeightedAdjacencyList [][]GraphEdge

func walk(graph WeightedAdjacencyList, current int, target int, seen []bool, path *[]int) bool {
	if seen[current] {
		return false
	}

	seen[current] = true

	*path = append(*path, current)

	if current == target {
		return true
	}

	list := graph[current]

	for i := 0; i < len(list); i++ {
		edge := list[i]
		if isEnd := walk(graph, edge.to, target, seen, path); isEnd {
			return true
		}
	}

	*path = (*path)[0 : len(*path)-1]

	return false
}

func DepthFirstSearch(graph WeightedAdjacencyList, source int, target int) []int {
	seen := make([]bool, len(graph))
	path := []int{}

	isEnd := walk(graph, source, target, seen, &path)

	if isEnd {
		return path
	}

	return []int{}
}
