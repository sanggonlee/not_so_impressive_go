package main

import (
	"fmt"
	"os"

	"github.com/sanggonlee/not_so_impressive_go_concurrency/asyncwriter"
)

func main() {
	f1, err := os.Create("foo")
	if err != nil {
		panic(err)
	}

	f2, err := os.Create("bar")
	if err != nil {
		panic(err)
	}

	f3, err := os.Create("baz")
	if err != nil {
		panic(err)
	}

	w := asyncwriter.AsyncMultiWriter(f1, f2, f3, os.Stdout)

	//fmt.Println(w.Close())
	fmt.Fprintln(w, "00000000001")
	// fmt.Println(w.Close())
	//n, err := fmt.Fprintln(w, "00000000002")
	n, err := w.Write([]byte("00000000002"))
	fmt.Println(n)
	fmt.Println(err)
}
