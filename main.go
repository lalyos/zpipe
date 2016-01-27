package main

import (
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func zip() {
	w := zlib.NewWriter(os.Stdout)
	io.Copy(w, os.Stdin)
	w.Close()
}

func unzip() {
	r, err := zlib.NewReader(os.Stdin)
	if err != nil {
		panic(err)
	}
	dat, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(dat))
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "-d" {
		unzip()
	} else {
		zip()
	}
}
