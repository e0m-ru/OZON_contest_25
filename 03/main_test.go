package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"sort"
	"strings"
	"testing"
)

func TestOzon04(t *testing.T) {
	var err error
	zipRC, _ := zip.OpenReader("even-strings.zip")
	defer zipRC.Close()
	for _, n := range zipRC.File { //len(zipRC.File)/2+1
		name := n.Name
		var inTestData, ExpTestDtat fs.File
		var expected []byte
		if !strings.Contains(name, ".a") {
			inTestData, err = zipRC.Open(n.Name)
			defer inTestData.Close()
			if err != nil {
				t.Error(err)
			}
			in = bufio.NewReaderSize(inTestData, 1024)
			ExpTestDtat, err = zipRC.Open(n.Name + ".a")
			if err != nil {
				t.Error(err)
			}
			expected, err = io.ReadAll(ExpTestDtat)
			if err != nil {
				t.Fatal(err)
			}
			defer ExpTestDtat.Close()
		} else {
			continue
		}
		var xxx bytes.Buffer
		out = bufio.NewWriter(&xxx)

		t.Run(fmt.Sprintf("Test %v", name), func(t *testing.T) {
			Ozon03()
			if string(expected) != string(xxx.Bytes()) {
				t.Errorf("\nEXPECTED:\n%v\nGOT\n%v\n",
					string(expected), string(xxx.Bytes()))
			}
		})
	}
}

func BenchmarkOzon03_b(b *testing.B) {

	var inTestData fs.File
	zipRC, _ := zip.OpenReader("even-strings.zip")
	defer zipRC.Close()
	out = bufio.NewWriter(io.Discard)

	var tests []string
	for _, f := range zipRC.File {
		name := f.Name
		if len(name) == 1 {
			name = "0" + name
		}

		if !strings.Contains(name, ".a") {
			tests = append(tests, name)
		} else {
			continue
		}
	}
	sort.Strings(tests)

	for _, n := range tests { //len(zipRC.File)/2+1
		name := n
		if n[0] == '0' {
			name = n[1:]
		}
		b.Run(fmt.Sprintf("test file %v", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				inTestData, _ = zipRC.Open(name)
				in = bufio.NewReader(inTestData)
				Ozon03()
			}
		})
		defer inTestData.Close()
	}
}

// go test -bench=. -benchmem -memprofile=mem.prof  -cpuprofile=cpu.prof
// go tool pprof -http=:8080 cpu.prof
func BenchmarkOzon03_a(b *testing.B) {

	out = bufio.NewWriter(io.Discard)
	in = bufio.NewReader(strings.NewReader(aaa))
	for i := 0; i < b.N; i++ {
		Ozon03()
	}

}
