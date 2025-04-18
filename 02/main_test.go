package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
	"testing"
)

func TestOzon_02(t *testing.T) {
	var err error
	zipRC, _ := zip.OpenReader("three-banks.zip")
	defer zipRC.Close()
	fmt.Printf("%v\n", len(zipRC.File))
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
			in = bufio.NewReader(inTestData)
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
			Ozon02()
			if string(expected) != string(xxx.Bytes()) {
				t.Errorf("\nEXPECTED:\n%v\nGOT\n%v\n",
					string(expected), string(xxx.Bytes()))
			}
		})
	}
}
