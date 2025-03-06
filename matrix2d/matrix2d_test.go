package matrix2d

import "testing"

var n = 10_000

func createMatrix(n int) *Matrix {
	rows, cols := n, n
	m := NewMatrix(rows, cols)
	k := 1
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			m.Set(i, j, k)
			k++
		}
	}
	return m
}

func BenchmarkMatrix_ContainsLinear(b *testing.B) {
	m := createMatrix(n)

	for i := 0; i <= n; i++ {
		m.ContainsLinear(i)
	}
}

func BenchmarkMatrix_ContainsExperimental(b *testing.B) {
	m := createMatrix(n)

	for i := 0; i <= n; i++ {
		m.ContainsExperimental(i, 0, 0, n, n)
	}
}

func BenchmarkMatrix_ContainsBS(b *testing.B) {
	m := createMatrix(n)

	for i := 0; i <= n; i++ {
		m.ContainsBS(i)
	}
}
