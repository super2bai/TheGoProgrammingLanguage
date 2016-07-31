package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"algorithm/bubblesort"
	"algorithm/qsort"
)

/**
Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating
another object (Reader or Writer) that also implements the interface but provides buffering
and some help for textual I/O.

Package flag implements command-line flag parsing.

Package fmt implements formatted I/O with functions analogous to C's printf and scanf.

Package io provides basic interfaces to I/O primitives.

Package os provides a platform-independent interface to operating system functionality.

Package strconv implements conversions to and from string representations of basic data types.
*/
var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile=", *outfile, "algorithm=", *algorithm)
	}

	values, err := readValus(*infile)
	if err == nil {
		tStart := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
		}
		tEnd := time.Now()
		fmt.Println("The soring process costs", tEnd.Sub(tStart), "to complete.")

		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}

func readValus(infile string) (values []int, err error) {
	file, errOpenFile := os.Open(infile)
	if errOpenFile != nil {
		fmt.Println("Failed to open the input file ", infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)

	for {
		line, isPrefix, errReadFile := br.ReadLine()
		if errReadFile != nil {
			if errReadFile != io.EOF {
				errReadFile = errOpenFile
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line,seems unexpected.")
			return
		}
		str := string(line) //转换字符数组为字符串

		value, errReadFile := strconv.Atoi(str)

		if errReadFile != nil {
			errOpenFile = errReadFile
			return
		}
		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)

	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}
