package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

type StringEntry struct {
	Odd, Even []byte
}

func main() {
	Ozon03Async()
}
func Ozon03() {
	var (
		n int // количество строк
		t int // количество наборов входных жанных
	)

	t = fastAtoi()
	for range t { // итерация по наборам
		n = fastAtoi()
		var ss = make([]StringEntry, n, n)
		for range n {
			line, _, _ := in.ReadLine()
			se, so := splitEvenOdd(line)
			ss = append(ss, StringEntry{Even: se, Odd: so})
		}
		out.WriteString(fmt.Sprintf("%d\n", compute(ss)))
		defer out.Flush()
	}
}

func compute(strings []StringEntry) (count int) {
	var n = len(strings)
	for i := range n {
		s := strings[i]
		for j := i + 1; j < n; j++ {
			t := strings[j]
			if compare(s, t) {
				count++
			}
		}
	}
	return
}

func compare(s, t StringEntry) bool {
	if bytes.Equal(s.Odd, t.Odd) && len(t.Odd) > 0 {
		return true
	} else if bytes.Equal(s.Even, t.Even) && len(t.Even) > 0 {
		return true
	}
	return false
}

func splitEvenOdd(data []byte) (even []byte, odd []byte) {
	for i, b := range data {
		if i%2 == 0 {
			even = append(even, b) // Чётные индексы
		} else {
			odd = append(odd, b) // Нечётные индексы
		}
	}
	return even, odd
}

func fastAtoi() (n int) {
	b, _ := in.ReadBytes('\n')
	b = bytes.TrimSpace(b)
	for _, ch := range b {
		n = n*10 + int(ch-'0')
	}
	return
}
