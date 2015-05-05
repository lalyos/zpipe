package main

import (
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	r, err := zlib.NewReader(os.Stdin)
	if err != nil {
		panic(err)
	}
	dat, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dat))
}
