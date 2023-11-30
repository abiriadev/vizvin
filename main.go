package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()

	fn := flag.Arg(0)

	if fn == "" {
		return
	}

	f, err := os.Open(fn)

	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)

	for i := 0; ; i++ {
		b, e := r.ReadByte()

		fmt.Printf("%02X ", b)

		if i%20 == 19 {
			fmt.Printf("\n")
		}

		if e == io.EOF {
			break
		}
	}
}
