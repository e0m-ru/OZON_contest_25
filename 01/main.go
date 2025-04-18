package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func Ozon01() {
	var t int    // количество наборов входных жанных
	var s string // строка
	fmt.Fscan(in, &t)

	for range t { // итерация по наборам
		fmt.Fscan(in, &s) // считываем строку
		assa(s)           // проверяем строку на соответствие условиям задачи
	}
	defer out.Flush() // сбрасываем буфер в файл
}

func main() {
	in = bufio.NewReader(strings.NewReader(aaa)) // для тестов
	Ozon01()                                     // запускаем основную функцию
}

func assa(s string) {
	// проверяем строку на соответствие условиям задачи
	var repeats int // количество повторов
	c := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != c {
			repeats++        // увеличиваем количество повторов
			if repeats > 1 { // если количество повторов больше 1, то строка не соответствует условиям задачи
				fmt.Fprintf(out, "%s\n", "NO")
				return
			}
		} else {
			repeats = 0
		}
	}
	if repeats == 0 { // если количество повторов равно 0, то строка соответствует условиям задачи
		fmt.Fprintf(out, "%s\n", "YES")
	} else {
		fmt.Fprintf(out, "%s\n", "NO")
	}
}
