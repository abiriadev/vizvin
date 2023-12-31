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

func resizePlot(src [][]int, srcSize int, destSize int) [][]int {
	dest := make([][]int, destSize)

	for i := range dest {
		dest[i] = make([]int, destSize)
	}

	rat := srcSize / destSize

	for i := 0; i < destSize; i++ {
		for j := 0; j < destSize; j++ {
			a := src[i*rat][j*rat]
			dest[i][j] = a
		}
	}

	return dest
}

func printPlot(plot [][]int, n int, scaleFunc Colorscale) {
	g := colorgrad.Turbo()

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			v := plot[i][j]
			c := g.At(scaleFunc.Scale(v))
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
	res := make([][]T, 256)

	for i := range res {
		res[i] = arr[i][:]
	}

	return res
}

func main() {
	bin := openFile()

	p := genplot(bin)

	destSize := 50

	dest := resizePlot(arr2slice(p), 256, destSize)

	printPlot(dest, destSize, NewDistinctScale(dest, destSize))
}
