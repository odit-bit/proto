package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/odit-bit/proto/resp"
)

func main() {

	s := "Set dollar 1000"

	b := resp.Encode(s)
	fmt.Printf("encoded %v", b)

	t, err := resp.Decode(bytes.NewReader(b))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("decoded", t.Array())
}
