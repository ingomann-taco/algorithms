package matrix1d

import (
	"algorithms/util"
	"errors"
	"fmt"
	"strings"
)

// Matrix Entries are stored in a 1d slice
type Matrix struct {
	rows, cols int
	entries    []int
}

// NewMatrix creates a new n x m matrix filled with zeros.
func NewMatrix(rows, cols int) *Matrix {
	return &Matrix{
		rows:    rows,
		cols:    cols,
		entries: make([]int, rows*cols),
	}
}

// Rows returns the number of rows in the matrix
func (m *Matrix) Rows() int {
	return m.rows
}

// Cols returns the number of columns in the matrix
func (m *Matrix) Cols() int {
	return m.cols
}

// FromSlice creates a new matrix, filling it from the given slice
func FromSlice(slice [][]int) *Matrix {
	rows := len(slice)
	cols := 0
	for _, row := range slice {
		if len(row) > cols {
			cols = len(row)
		}
	}

	m := NewMatrix(rows, cols)

	for i, row := range slice {
		for j, val := range row {
			m.Set(i, j, val)
		}
	}

	return m
}

// index calculates the slice index for given row and column
func (m *Matrix) index(row, col int) (int, error) {
	k := row*m.cols + col
	if k >= len(m.entries) {
		return 0, fmt.Errorf("index out of bounds")
	}
	return k, nil
}

// Set sets the value at (row, col) to val
func (m *Matrix) Set(row, col int, val int) error {
	k, err := m.index(row, col)
	if err != nil {
		return err
	}
	m.entries[k] = val
	return nil
}

// Get retrieves the value at (row, col)
func (m *Matrix) Get(row, col int) (int, error) {
	k, err := m.index(row, col)
	if err != nil {
		return 0, err
	}
	return m.entries[k], nil
}

// Row returns a row as a slice
func (m *Matrix) Row(row int) ([]int, error) {
	if row < 0 || m.rows <= row {
		return nil, errors.New("row index out of bounds")
	}
	start := row * m.cols
	return m.entries[start : start+m.cols], nil
}

// ContainsLinear Checks if the target exists in the matrix in linear time, constant memory
// If found, returns true and the row and column indices
func (m *Matrix) ContainsLinear(target int) (bool, int, int) {
	i, j := 0, m.Cols()-1

	for i < m.Rows() && 0 <= j {
		el, _ := m.Get(i, j)
		if el == target {
			return true, i, j
		} else if el > target {
			j--
		} else {
			i++
		}
	}

	return false, 0, 0
}

// ContainsExperimental checks if the target exists in the matrix using a divide and conquer approach.
// The search is restricted to the block described by (ai, aj) - (bi, bj).
//
// Parameters:
//
//	ai, aj - starting row and column indices of the search window (half-open interval)
//	bi, bj - ending row and column indices of the search window (half-open interval)
//	target - the value to search for
//
// Returns:
//
//	bool - true if the target is found, false otherwise
//	int  - row index of the target if found
//	int  - column index of the target if found
//	error - error if the search window is invalid
func (m *Matrix) ContainsExperimental(ai, aj, bi, bj, target int) (bool, int, int, error) {
	if err := m.validateWindow(ai, aj, bi, bj); err != nil {
		return false, 0, 0, err
	}

	rows, cols := bi-ai, bj-aj
	if rows == 0 || cols == 0 {
		return false, 0, 0, nil
	} else if rows == 1 && cols == 1 {
		v, _ := m.Get(ai, aj)
		return v == target, ai, aj, nil
	}

	mMin, _ := m.Get(ai, aj)
	mMax, _ := m.Get(bi-1, bj-1)
	if target < mMin || mMax < target {
		return false, 0, 0, nil
	}

	rowMid, colMid := ai+(rows/2), aj+(cols/2)

	// Top-left quadrant
	if found, row, col, _ := m.ContainsExperimental(ai, aj, rowMid, colMid, target); found {
		return true, row, col, nil
	}
	// Top-right quadrant
	if found, row, col, _ := m.ContainsExperimental(ai, colMid, rowMid, bj, target); found {
		return true, row, col, nil
	}
	// Bottom-left quadrant
	if found, row, col, _ := m.ContainsExperimental(rowMid, aj, bi, colMid, target); found {
		return true, row, col, nil
	}
	// Bottom-right quadrant
	if found, row, col, _ := m.ContainsExperimental(rowMid, colMid, bi, bj, target); found {
		return true, row, col, nil
	}

	return false, 0, 0, nil
}

// validateWindow checks if the window is valid for the given matrix
func (m *Matrix) validateWindow(ai, aj, bi, bj int) error {
	if ai < 0 || m.rows < ai {
		return fmt.Errorf("invalid ai %d", ai)
	} else if bi < 0 || m.rows < bi {
		return fmt.Errorf("invalid bi %d", bi)
	} else if aj < 0 || m.cols < aj {
		return fmt.Errorf("invalid aj %d", aj)
	} else if bj < 0 || m.cols < bj {
		return fmt.Errorf("invalid bj %d", bj)
	} else if bi < ai || bj < aj {
		return fmt.Errorf("invalid window (%d, %d) - (%d, %d)", ai, aj, bi, bj)
	}
	return nil
}

// ContainsBS Checks if the target exists in the matrix using binary search in n*log(n) time, constant memory
// If found, returns true and the row and column indices
func (m *Matrix) ContainsBS(target int) (bool, int, int) {
	for i := 0; i < m.Rows(); i++ {
		row, _ := m.Row(i)
		col, found := util.BinarySearch(row, target)
		if found {
			return found, i, col
		}
	}
	return false, 0, 0
}

// String displays the matrix.
func (m *Matrix) String() string {
	var maxWidth int
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			val, _ := m.Get(i, j)
			width := len(fmt.Sprintf("%d", val))
			if width > maxWidth {
				maxWidth = width
			}
		}
	}

	var sb strings.Builder
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			val, _ := m.Get(i, j)
			sb.WriteString(fmt.Sprintf("%*d ", maxWidth, val))
		}
		if i != m.rows-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
