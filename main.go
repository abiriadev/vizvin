package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/abiriadev/iris"
	"github.com/mazznoer/colorgrad"
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

func plotMax(plot [][]int, n int) int {
	mx := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if plot[i][j] > mx {
				mx = plot[i][j]
			}
		}
	}

	return mx
}

func printPlot(plot [][]int, n int) {
	mx := plotMax(plot, n)
	g := colorgrad.Turbo()

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			v := plot[i][j]
			c := g.At(float64(v) / float64(mx))
			fmt.Printf("%s  ", iris.ColorBg(c))
		}
		fmt.Printf("%s\n", iris.Reset)
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

func arr2slice[T any](arr [256][256]T) [][]T {
	res := make([][]T, 256*256)

	for i := 0; i < 256; i++ {
		res[i] = arr[i][:]
	}

	return res
}

func main() {
	bin := openFile()

	p := genplot(bin)

	printPlot(arr2slice(p), 256)
}
