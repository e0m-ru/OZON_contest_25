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

type Server struct {
	num int
	mms int
}

type File struct {
	num int
	mem int
	srv int
}

func (f *File) calcTimes(servers *[]Server) {
	for _, s := range *servers {
		fmt.Printf("%d ", (f.mem+s.mms-1)/s.mms)
	}
	fmt.Printf("\n")
}

func Ozon05() {
	var servers []Server // массив серверов
	var files []File     // массив изображений

	var a int // количество наборов входных жанных
	fmt.Fscan(in, &a)

	for range a { // итерация по наборам
		var numServers int // количество серверов
		fmt.Fscan(in, &numServers)

		servers = make([]Server, numServers)

		for i := range numServers { // итерация по серверам
			var ms int // пропускная способность сервера №n
			fmt.Fscan(in, &ms)
			servers[i] = Server{
				num: i,
				mms: ms,
			}
		}

		var numImages int // количество изображений
		fmt.Fscan(in, &numImages)

		files = make([]File, numImages)
		for i := range numImages { // итерация по изображениям
			var im int // вес изображения m
			fmt.Fscan(in, &im)
			files[i] = File{
				num: i,
				mem: im,
			}
		}

		calculate(files, &servers)
	}

	defer out.Flush()
}

func main() {
	in = bufio.NewReader(strings.NewReader(bbb))
	Ozon05()
}

func calculate(files []File, servers *[]Server) {
	// L, R := 0, 1
	for _, f := range files {
		f.calcTimes(servers)
	}
}
