package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"sync"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

type StringEntry struct {
	Odd, Even         []byte
	HashOdd, HashEven [32]byte
}

func main() {
	Ozon03Async()
}
func Ozon03Async() {
	var (
		wg sync.WaitGroup
		t  int
		n  int
	)

	t = fastAtoi()

	results := make(chan struct{ n, c int }, t)
	for q := range t { // итерация по наборам
		n = fastAtoi()
		ss := make([]StringEntry, n)
		for i := range n {
			line, _, _ := in.ReadLine()
			se, so := splitEvenOdd(line)
			ss[i] = StringEntry{HashEven: getHash(se), HashOdd: getHash(so), Even: se, Odd: so}
		}
		wg.Add(1)
		go func(wg *sync.WaitGroup, res chan struct{ n, c int }, ss []StringEntry) {
			results <- struct{ n, c int }{n: q, c: compute(ss)}
			wg.Done()
		}(&wg, results, ss)
	}
	wg.Wait()
	close(results)
	aaa := make([]int, t)
	for r := range results {
		aaa[r.n] = r.c
	}
	for _, w := range aaa {
		out.WriteString(fmt.Sprintf("%d\n", w))
	}
	defer out.Flush()
}

func compute(strings []StringEntry) (count int) {
	l := len(strings)
	for i := range l {
		for j := i + 1; j < l; j++ {
			if compare(&strings[i], &strings[j]) {
				count++
			}
		}
	}
	return
}

func compare(s, t *StringEntry) bool {
	if (s.HashEven == t.HashEven) && ((len(s.Even) > 0) && (len(t.Even) > 0)) {
		return true
	} else if (s.HashOdd == t.HashOdd) && ((len(s.Odd) > 0) && (len(t.Odd) > 0)) {
		return true
	}
	return false
}

func splitEvenOdd(data []byte) (even []byte, odd []byte) {
	for i, b := range data {
		if i%2 == 0 {
			even = append(even, b)
		} else {
			odd = append(odd, b)
		}
	}
	return even, odd
}

func fastAtoi() (n int) {
	b, _ := in.ReadBytes('\n')
	for _, ch := range b[:len(b)-1] {
		n = n*10 + int(ch-'0')
	}
	return
}

func getHash(data []byte) [32]byte {
	return sha256.Sum256(data)
}
