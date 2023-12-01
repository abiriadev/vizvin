package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func genplot(r bufio.Reader) [256][256]int {
	var plot [256][256]int

	last, e := r.ReadByte()

	if e == io.EOF {
		return plot
	}

	for {
		curr, e := r.ReadByte()

		if e == io.EOF {
			break
		}

		fmt.Printf("%02x, %02x\n", last, curr)

		last = curr
	}

	return plot
}

func openFile() bufio.Reader {
	flag.Parse()

	fn := flag.Arg(0)

	if fn == "" {
		panic("pass binary file to visualize")
	}

	f, err := os.Open(fn)

	if err != nil {
		panic(err)
	}

	return *bufio.NewReader(f)
}

func main() {
	bin := openFile()

	_ = genplot(bin)
}
