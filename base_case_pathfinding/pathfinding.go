package basecasepathfinding

type Point struct {
	X int
	Y int
}

type Square string

const (
	SquareEmpty Square = ""
	SquareWall  Square = "#"
	SquareEnd   Square = "E"
)

var directions = []Point{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func walk(maze [][]Square, seen [][]bool, path []Point, current Point) (bool, []Point) {
	if current.X < 0 || current.X >= len(maze[0]) || current.Y < 0 || current.Y >= len(maze) {
		return false, path
	}

	if maze[current.Y][current.X] == SquareWall {
		return false, path
	}

	if seen[current.Y][current.X] {
		return false, path
	}

	path = append(path, current)

	if maze[current.Y][current.X] == SquareEnd {
		return true, path
	}

	seen[current.Y][current.X] = true

	for i := 0; i < len(directions); i++ {
		if isEnd, path := walk(maze, seen, path, Point{current.X + directions[i].X, current.Y + directions[i].Y}); isEnd {
			return true, path
		}
	}

	path = path[:len(path)-1]
	return false, path
}

func Solve(maze [][]Square, start Point) (bool, []Point) {
	seen := make([][]bool, len(maze))
	path := make([]Point, 0)

	for i := range maze {
		seen[i] = make([]bool, len(maze[0]))
	}

	return walk(maze, seen, path, start)
}
