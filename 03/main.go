package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sync"
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
func Ozon03Async() {
	var (
		n  int // количество строк
		t  int // количество наборов входных жанных
		wg sync.WaitGroup
	)

	t = fastAtoi()
	results := make(chan struct{ n, c int }, t)
	for q := range t { // итерация по наборам

		n = fastAtoi()
		var ss = make([]StringEntry, n)
		for i := range n {
			line, _, _ := in.ReadLine()
			se, so := splitEvenOdd(line)
			ss[i] = StringEntry{Even: se, Odd: so}
		}
		wg.Add(1)
		go func(wg *sync.WaitGroup, res chan struct{ n, c int }, ss []StringEntry) {
			results <- struct{ n, c int }{n: q, c: compute(ss)}
			wg.Done()
		}(&wg, results, ss)
	}
	wg.Wait()
	close(results)
	var aaa = make([]int, t)
	for r := range results {
		aaa[r.n] = r.c
	}
	for _, w := range aaa {
		out.WriteString(fmt.Sprintf("%d\n", w))
	}

	defer out.Flush()
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
	if len(s.Even) != len(t.Even) && len(s.Odd) != len(t.Odd) {
		return false
	}
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
