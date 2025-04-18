package main

import (
	"fmt"
	"runtime"
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

	for range runtime.NumCPU() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasks {
				if compare(task.s, task.t) {
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
