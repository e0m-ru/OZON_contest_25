package main

import (
	"bufio"
	"fmt"
	"hash/maphash"
	"os"
	"sync"
)

var (
	in   = bufio.NewReader(os.Stdin)
	out  = bufio.NewWriter(os.Stdout)
	line []byte
)

type StringEntry struct {
	HashOdd, HashEven [32]byte
}

var (
	hasherEven maphash.Hash
	hasherOdd  maphash.Hash
)

func readLinesRaw(reader *bufio.Reader) [][][]uint64 {
	line, _ = reader.ReadSlice('\n')
	t := fastAtoi(line)
	allTesets := make([][][]uint64, t)
	for i := range t {
		line, _ = reader.ReadSlice('\n')
		n := fastAtoi(line)
		lines := make([][]uint64, n)
		for q := range n {
			line, _ = reader.ReadSlice('\n')
			lines[q] = computeHashPair(line[:len(line)-1])
		}
		allTesets[i] = lines
	}
	return allTesets
}

func main() {
	seed := maphash.MakeSeed()
	hasherEven.SetSeed(seed)
	hasherOdd.SetSeed(seed)
	Ozon03(readLinesRaw(in))
}

func Ozon03(data [][][]uint64) {
	var (
		wg sync.WaitGroup
	)
	results := make(chan struct{ n, c int }, len(data))
	for i, q := range data { // итерация по наборам
		wg.Add(1)
		go func(wg *sync.WaitGroup, res chan struct{ n, c int }, ss [][]uint64) {
			results <- struct{ n, c int }{n: i, c: compute(ss)}
			wg.Done()
		}(&wg, results, q)
	}
	wg.Wait()
	close(results)
	aaa := make([]int, len(data))
	for r := range results {
		aaa[r.n] = r.c
	}
	for _, w := range aaa {
		out.WriteString(fmt.Sprintf("%d\n", w))
	}
	defer out.Flush()
}

func compute(strings [][]uint64) (count int) {
	l := len(strings)
	for i := range l {
		for j := i + 1; j < l; j++ {
			if compare(strings[i], strings[j]) {
				count++
			}
		}
	}
	return
}

func compare(a, b []uint64) bool {

	if ((a[0] == b[0]) && (b[0] != 0)) || ((a[1] == b[1]) && (b[1] != 0)) {
		return true
	}

	return false
}

func computeHashPair(data []byte) (res []uint64) {
	hasherEven.Reset()
	hasherOdd.Reset()
	var f2 = false
	for i, b := range data {
		if i%2 == 0 {
			hasherEven.WriteByte(b)
		} else {
			hasherOdd.WriteByte(b)
			f2 = true
		}
	}
	if f2 {
		return []uint64{hasherEven.Sum64(), hasherOdd.Sum64()}
	} else {
		return []uint64{hasherEven.Sum64(), 0}
	}
}

func fastAtoi(b []byte) (n int) {
	for _, ch := range b[:len(b)-1] {
		n = n*10 + int(ch-'0')
	}
	return
}
