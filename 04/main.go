package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func Ozon04() {
	var a int    // количество наборов входных жанных
	var k int    //необходимая для победы длина строки, столбца или диагонали
	var n, m int // размеры поля

	fmt.Fscan(in, &a)
	for range a { // итерация по наборам
		fmt.Fscan(in, &k)
		fmt.Fscan(in, &n, &m)
		var s = make([][]rune, n) // поле

		for i := range n {
			var x string
			fmt.Fscan(in, &x)
			s[i] = splitString(m, k, x)
			if s[i] == nil {
				break
			}
		}
		f := false
		for _, v := range getRows(s, k) {
			if stringProcessor(v, k) {
				out.WriteString("YES\n")
				f = true
				break
			}
		}
		for _, v := range getColumns(s, k) {
			if stringProcessor(v, k) {
				out.WriteString("YES\n")
				f = true
				break
			}
		}
		for _, v := range getDiagonals(s, k) {
			if stringProcessor(v, k) {
				out.WriteString("YES\n")
				f = true
				break
			}
		}
		if !f {
			out.WriteString("NO\n")
		}
	}

	defer out.Flush()
}

func splitString(m, k int, s string) []rune {
	row := make([]rune, m)
	for i, c := range s {
		row[i] = c
	}
	return row
}

func stringProcessor(s []rune, k int) (f bool) {
	if len(s) < k {
		return false
	}

	xCounter := 0 // счетчик для X
	g := false
	for _, c := range s {
		switch c {
		case 'X':
			xCounter++
			if xCounter == k && !g {
				return false
			} else if xCounter == k && !g {
				return true
			}
		case 'O':
			xCounter = 0
			g = false
		case '.':
			if g {
				xCounter = 0
				g = false
			} else {
				xCounter++
				g = true
			}
		}
	}
	if xCounter == k && g {
		return true
	} else {
		return false
	}

}

func main() {
	Ozon04()
}

func getRows(grid [][]rune, k int) [][]rune {
	var rows [][]rune
	for _, row := range grid {
		if len(row) >= k {
			rows = append(rows, row)
		}
	}
	return rows
}

func getColumns(grid [][]rune, k int) [][]rune {
	var columns [][]rune
	m := len(grid[0]) // количество колонок
	for j := 0; j < m; j++ {
		var column []rune
		for i := 0; i < len(grid); i++ {
			column = append(column, grid[i][j])
		}
		if len(column) >= k {
			columns = append(columns, column)
		}
	}
	return columns
}

func getDiagonals(grid [][]rune, k int) [][]rune {
	var diagonals [][]rune
	n, m := len(grid), len(grid[0])

	// Главные диагонали (слева направо)
	for start := 0; start < n; start++ {
		var diagonal []rune
		for i, j := start, 0; i < n && j < m; i, j = i+1, j+1 {
			diagonal = append(diagonal, grid[i][j])
		}
		if len(diagonal) >= k {
			diagonals = append(diagonals, diagonal)
		}
	}
	for start := 1; start < m; start++ {
		var diagonal []rune
		for i, j := 0, start; i < n && j < m; i, j = i+1, j+1 {
			diagonal = append(diagonal, grid[i][j])
		}
		if len(diagonal) >= k {
			diagonals = append(diagonals, diagonal)
		}
	}

	// Побочные диагонали (справа налево)
	for start := 0; start < n; start++ {
		var diagonal []rune
		for i, j := start, m-1; i < n && j >= 0; i, j = i+1, j-1 {
			diagonal = append(diagonal, grid[i][j])
		}
		if len(diagonal) >= k {
			diagonals = append(diagonals, diagonal)
		}
	}
	for start := m - 2; start >= 0; start-- {
		var diagonal []rune
		for i, j := 0, start; i < n && j >= 0; i, j = i+1, j-1 {
			diagonal = append(diagonal, grid[i][j])
		}
		if len(diagonal) >= k {
			diagonals = append(diagonals, diagonal)
		}
	}

	return diagonals
}
