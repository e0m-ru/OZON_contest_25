package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestOzon05(t *testing.T) {
	zipRC, _ := zip.OpenReader("content-delivery.zip")
	defer zipRC.Close()

	for i := 1; i < 5; i++ { //len(zipRC.File)/2+1
		inTestData, _ := zipRC.Open(fmt.Sprint(i))
		ExpTestDtat, _ := zipRC.Open(fmt.Sprint(i) + ".a")
		defer ExpTestDtat.Close()
		in = bufio.NewReader(inTestData)
		var xxx bytes.Buffer
		out = bufio.NewWriter(&xxx)
		expected, _ := io.ReadAll(ExpTestDtat)
		t.Run(fmt.Sprintf("Test№ %v", i), func(t *testing.T) {
			Ozon05()
			if string(expected) != string(xxx.Bytes()) {
				t.Errorf("\nEXPECTED:\n%v\nGOT\n%v\n",
					string(expected), string(xxx.Bytes()))
			}
		})
	}
}
