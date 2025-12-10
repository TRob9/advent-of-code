package grid

import (
	"fmt"
	"strings"
)

// Grid[T] is a simple row-major 2D grid of values of type T.
type Grid[T any] struct {
	cols, rows int
	data       []T // length == cols * rows
}

// New creates a cols x rows grid and fills with the zero value of T.
func New[T any](cols, rows int) *Grid[T] {
	if cols <= 0 || rows <= 0 {
		panic("cols and rows must be > 0")
	}
	return &Grid[T]{cols: cols, rows: rows, data: make([]T, cols*rows)}
}

// From2D constructs a Grid from a 2D slice. All inner slices must have same length.
func From2D[T any](arr [][]T) *Grid[T] {
	if len(arr) == 0 || len(arr[0]) == 0 {
		panic("empty 2D slice")
	}
	rows := len(arr)
	cols := len(arr[0])
	g := New[T](cols, rows)
	for y := 0; y < rows; y++ {
		if len(arr[y]) != cols {
			panic("ragged 2D slice")
		}
		for x := 0; x < cols; x++ {
			g.Set(x, y, arr[y][x])
		}
	}
	return g
}

// Size returns (cols, rows)
func (g *Grid[T]) Size() (int, int) { return g.cols, g.rows }

// Index converts x,y to internal index (row-major).
// Caller must ensure InBounds before calling Index.
func (g *Grid[T]) Index(x, y int) int { return y*g.cols + x }

// InBounds checks whether x,y are within the grid.
func (g *Grid[T]) InBounds(x, y int) bool {
	return x >= 0 && x < g.cols && y >= 0 && y < g.rows
}

// Get returns value at x,y and true if in bounds, otherwise zero value and false.
func (g *Grid[T]) Get(x, y int) (T, bool) {
	var zero T
	if !g.InBounds(x, y) {
		return zero, false
	}
	return g.data[g.Index(x, y)], true
}

// MustGet returns the value at x,y or panics if out-of-bounds.
func (g *Grid[T]) MustGet(x, y int) T {
	if !g.InBounds(x, y) {
		panic(fmt.Sprintf("Get: out of bounds (%d,%d)", x, y))
	}
	return g.data[g.Index(x, y)]
}

// Set sets value at x,y. Returns false if out-of-bounds.
func (g *Grid[T]) Set(x, y int, v T) bool {
	if !g.InBounds(x, y) {
		return false
	}
	g.data[g.Index(x, y)] = v
	return true
}

// ForEach calls fn(x,y,val) for every cell in row-major order.
// If fn returns false, iteration stops early.
func (g *Grid[T]) ForEach(fn func(x, y int, v T) bool) {
	for y := 0; y < g.rows; y++ {
		for x := 0; x < g.cols; x++ {
			if !fn(x, y, g.data[g.Index(x, y)]) {
				return
			}
		}
	}
}

// Neighbors returns coordinates of orthogonal neighbors of (x,y).
// If includeDiag is true, diagonal neighbors are included too.
func (g *Grid[T]) Neighbors(x, y int, includeDiag bool) [][2]int {
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	if includeDiag {
		dirs = append(dirs, [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}}...)
	}
	out := make([][2]int, 0, len(dirs))
	for _, d := range dirs {
		nx, ny := x+d[0], y+d[1]
		if g.InBounds(nx, ny) {
			out = append(out, [2]int{nx, ny})
		}
	}
	return out
}

// String prints the grid rows with fmt.Sprint of each element separated by sep.
// Use for debugging / quick view.
func (g *Grid[T]) String(sep string) string {
	var sb strings.Builder
	for y := 0; y < g.rows; y++ {
		for x := 0; x < g.cols; x++ {
			if x > 0 {
				sb.WriteString(sep)
			}
			sb.WriteString(fmt.Sprint(g.data[g.Index(x, y)]))
		}
		if y < g.rows-1 {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}
