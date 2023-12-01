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

		plot[last][curr]++

		last = curr
	}

	return plot
}

func printPlot(plot [256][256]int) {
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			fmt.Printf("%d ", plot[i][j])
		}
		fmt.Printf("\n")
	}
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

	p := genplot(bin)

	printPlot(p)
}
