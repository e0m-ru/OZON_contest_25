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

var (
	n int // количество строк
	t int // количество наборов входных жанных
)

type StringEntry struct {
	Odd  []byte // строка
	Even []byte // строка
}

func Ozon03() {
	fmt.Fscan(in, &t)
	in.ReadByte()
	for range t { // итерация по наборам
		fmt.Fscan(in, &n)
		in.ReadByte()
		var ss = make([]StringEntry, n)
		for i := range n {
			line, _, _ := in.ReadLine()
			se, so := splitEvenOdd(line)
			ss[i] = StringEntry{Even: se, Odd: so}
		}
		assa(ss, n)
	}
	defer out.Flush()
}

func main() {
	Ozon03()
}
func assa(strings []StringEntry, n int) {
	var count int
	for i := range n {
		s := strings[i]
		for j := i + 1; j < n; j++ {
			t := strings[j]
			if bytes.Equal(s.Odd, t.Odd) && len(t.Odd) > 0 {
				count++
			} else if bytes.Equal(s.Even, t.Even) && len(t.Even) > 0 {
				count++
			}
		}
	}
	out.WriteString(fmt.Sprintf("%d\n", count))
}

const (
	Even = iota
	Odd
)

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
