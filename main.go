package main

import (
	"algorithms/matrix1d"
	"fmt"
	"os"
)

func main() {
	rows, cols := 10, 10
	m := matrix1d.NewMatrix(rows, cols)
	k := 1
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			m.Set(i, j, k)
			k++
		}
	}

	fmt.Println("Matrix:")
	fmt.Println(m)
	fmt.Println()

	targets := []int{57, 271}

	doLinearSearch(m, targets)
	fmt.Println()
	doBSSearch(m, targets)
	fmt.Println()
	doExperimentalSearch(m, targets)
}

func doLinearSearch(m *matrix1d.Matrix, targets []int) {
	fmt.Println("Linear time search algorithm")
	for _, target := range targets {
		found, row, col := m.ContainsLinear(target)
		if found {
			fmt.Printf("- Found %d at (%d, %d)\n", target, row, col)
		} else {
			fmt.Printf("- %d not found\n", target)
		}
	}
}

func doBSSearch(m *matrix1d.Matrix, targets []int) {
	fmt.Println("Binary Search algorithm")
	for _, target := range targets {
		found, row, col := m.ContainsBS(target)
		if found {
			fmt.Printf("- Found %d at (%d, %d)\n", target, row, col)
		} else {
			fmt.Printf("- %d not found\n", target)
		}
	}
}

func doExperimentalSearch(m *matrix1d.Matrix, targets []int) {
	fmt.Println("Experimental search algorithm")
	for _, target := range targets {
		found, row, col, err := m.ContainsExperimental(0, 0, m.Rows(), m.Cols(), target)
		if err != nil {
			fmt.Fprintf(os.Stderr, "experimental search: %v", err)
			continue
		}
		if found {
			fmt.Printf("- Found %d at (%d, %d)\n", target, row, col)
		} else {
			fmt.Printf("- %d not found\n", target)
		}
	}
}
