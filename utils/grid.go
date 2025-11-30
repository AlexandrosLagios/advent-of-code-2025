package utils

// Point represents a 2D point
type Point struct {
	X, Y int
}

// Direction represents a 2D direction
type Direction struct {
	DX, DY int
}

// Common directions
var (
	North = Direction{0, -1}
	South = Direction{0, 1}
	East  = Direction{1, 0}
	West  = Direction{-1, 0}

	NorthEast = Direction{1, -1}
	NorthWest = Direction{-1, -1}
	SouthEast = Direction{1, 1}
	SouthWest = Direction{-1, 1}
)

// Cardinal4 returns the 4 cardinal directions
func Cardinal4() []Direction {
	return []Direction{North, East, South, West}
}

// Cardinal8 returns the 8 directions (including diagonals)
func Cardinal8() []Direction {
	return []Direction{North, NorthEast, East, SouthEast, South, SouthWest, West, NorthWest}
}

// Move returns a new point moved in the given direction
func (p Point) Move(d Direction) Point {
	return Point{p.X + d.DX, p.Y + d.DY}
}

// MoveBy returns a new point moved by dx, dy
func (p Point) MoveBy(dx, dy int) Point {
	return Point{p.X + dx, p.Y + dy}
}

// Neighbors4 returns the 4 cardinal neighbors
func (p Point) Neighbors4() []Point {
	return []Point{
		{p.X, p.Y - 1}, // North
		{p.X + 1, p.Y}, // East
		{p.X, p.Y + 1}, // South
		{p.X - 1, p.Y}, // West
	}
}

// Neighbors8 returns the 8 neighbors (including diagonals)
func (p Point) Neighbors8() []Point {
	return []Point{
		{p.X, p.Y - 1},     // North
		{p.X + 1, p.Y - 1}, // NorthEast
		{p.X + 1, p.Y},     // East
		{p.X + 1, p.Y + 1}, // SouthEast
		{p.X, p.Y + 1},     // South
		{p.X - 1, p.Y + 1}, // SouthWest
		{p.X - 1, p.Y},     // West
		{p.X - 1, p.Y - 1}, // NorthWest
	}
}

// InBounds checks if a point is within grid bounds
func (p Point) InBounds(width, height int) bool {
	return p.X >= 0 && p.X < width && p.Y >= 0 && p.Y < height
}

// Grid utility functions

// InBoundsInt checks if coordinates are within grid bounds
func InBoundsInt(x, y, width, height int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

// TransposeGrid transposes a 2D grid
func TransposeGrid[T any](grid [][]T) [][]T {
	if len(grid) == 0 {
		return grid
	}
	rows := len(grid)
	cols := len(grid[0])
	transposed := make([][]T, cols)
	for i := range transposed {
		transposed[i] = make([]T, rows)
		for j := range transposed[i] {
			transposed[i][j] = grid[j][i]
		}
	}
	return transposed
}

// RotateGridCW rotates a 2D grid 90 degrees clockwise
func RotateGridCW[T any](grid [][]T) [][]T {
	if len(grid) == 0 {
		return grid
	}
	rows := len(grid)
	cols := len(grid[0])
	rotated := make([][]T, cols)
	for i := range rotated {
		rotated[i] = make([]T, rows)
		for j := range rotated[i] {
			rotated[i][j] = grid[rows-1-j][i]
		}
	}
	return rotated
}

// FlipGridHorizontal flips a grid horizontally
func FlipGridHorizontal[T any](grid [][]T) [][]T {
	flipped := make([][]T, len(grid))
	for i, row := range grid {
		flipped[i] = make([]T, len(row))
		for j := range row {
			flipped[i][j] = row[len(row)-1-j]
		}
	}
	return flipped
}
