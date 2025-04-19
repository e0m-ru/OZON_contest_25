package main

import (
	"fmt"
	"sync"
)

func computeAsync(strings []StringEntry) (count int) {
	type task struct {
		s, t StringEntry
	}
	var n = len(strings)
	var wg sync.WaitGroup
	var mu sync.Mutex
	tasks := make(chan task)

	for range n {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasks {
				if compare(&task.s, &task.t) {
					mu.Lock()
					count++
					mu.Unlock()
				}
			}
		}()
	}

	for i := range n {
		s := strings[i]
		for j := i + 1; j < n; j++ {
			t := strings[j]
			tasks <- task{s, t}
		}
	}
	close(tasks)
	wg.Wait()
	return
}

func Ozon03() {
	var (
		n int // количество строк
		t int // количество наборов входных жанных
	)

	t = fastAtoi()
	// fmt.Fscan(in, &t)

	for range t { // итерация по наборам
		n = fastAtoi()
		// fmt.Fscan(in, &n)

		var ss = make([]StringEntry, n)
		for range n {
			line, _, _ := in.ReadLine()
			se, so := splitEvenOdd(line)
			ss = append(ss, StringEntry{Even: se, Odd: so})
		}
		out.WriteString(fmt.Sprintf("%d\n", compute(ss)))
		defer out.Flush()
	}
}
