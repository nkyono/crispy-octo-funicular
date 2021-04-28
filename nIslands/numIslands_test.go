package test

import (
	"fmt"
	"testing"
)

type index struct {
	i int
	j int
}

// tried to reduce memory usage by removing visited array
// Didn't help, pretty sure the stack is killing the memory usage
func NumIslandsMem(grid [][]byte) int {
	res := 0

	stack := []index{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				stack = append(stack, index{i, j})
				res++
			}
			for len(stack) > 0 {
				v := stack[0]
				if !(v.i < 0 || v.j < 0 || v.i >= len(grid) || v.j >= len(grid[0]) || grid[v.i][v.j] == '0') {
					stack = append(stack, index{v.i - 1, v.j}, index{v.i, v.j - 1}, index{v.i + 1, v.j}, index{v.i, v.j + 1})
					grid[v.i][v.j] = '0'
				}
				stack = stack[1:]
			}
		}
	}

	return res
}

func solve(grid [][]byte, i, j int, visited [][]bool) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || visited[i][j] || grid[i][j] == '0' {
		return
	}
	visited[i][j] = true
	solve(grid, i-1, j, visited)
	solve(grid, i+1, j, visited)
	solve(grid, i, j-1, visited)
	solve(grid, i, j+1, visited)
}

func NumIslandsRecursive(grid [][]byte) int {
	res := 0
	visited := make([][]bool, len(grid))
	for i, v := range grid {
		visited[i] = make([]bool, len(v))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				solve(grid, i, j, visited)
				res++
			}
		}
	}

	return res
}

func NumIslandsIterative(grid [][]byte) int {
	res := 0
	visited := make([][]bool, len(grid))
	for i, v := range grid {
		visited[i] = make([]bool, len(v))
	}

	stack := []index{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				stack = append(stack, index{i, j})
				res++
			}
			for len(stack) > 0 {
				v := stack[0]
				if !(v.i < 0 || v.j < 0 || v.i >= len(grid) || v.j >= len(grid[0]) || visited[v.i][v.j] || grid[v.i][v.j] == '0') {
					stack = append(stack, index{v.i - 1, v.j}, index{v.i, v.j - 1}, index{v.i + 1, v.j}, index{v.i, v.j + 1})
					visited[v.i][v.j] = true
				}
				stack = stack[1:]
			}
		}
	}

	return res
}

func TestNumIslands(t *testing.T) {
	var tests = []struct {
		in   [][]byte
		want int
	}{
		{[][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}, 1},
		{[][]byte{
			{'1', '0'},
			{'0', '1'},
		}, 2},
		{[][]byte{
			{'1', '1'},
			{'1', '1'},
		}, 1},
		{[][]byte{
			{'0', '0'},
			{'0', '0'},
		}, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Iterative: %v", tt.in)
		t.Run(testname, func(t *testing.T) {
			ans := NumIslandsIterative(tt.in)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Recursive: %v", tt.in)
		t.Run(testname, func(t *testing.T) {
			ans := NumIslandsRecursive(tt.in)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Mem: %v", tt.in)
		t.Run(testname, func(t *testing.T) {
			ans := NumIslandsMem(tt.in)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
