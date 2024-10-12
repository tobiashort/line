package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func usage() {
	fmt.Print(`Usage: line -n NUM [FILE]
Reads specified line number from FILE or STDIN if FILE not given.
`)
}

func main() {
	lineNum := 0

	flag.Usage = usage
	flag.IntVar(&lineNum, "n", 0, "the line number")
	flag.Parse()

	if lineNum < 1 {
		usage()
		os.Exit(1)
	}

	var file *os.File
	defer file.Close()

	if flag.NArg() > 1 {
		usage()
		os.Exit(1)
	} else if flag.NArg() == 1 {
		fileName := flag.Arg(0)
		var err error
		file, err = os.Open(fileName)
		if err != nil {
			panic(err)
		}
	} else {
		file = os.Stdin
	}

	reader := bufio.NewReader(file)

	line := ""
	currentLineNum := 0
	for currentLineNum < lineNum {
		var err error
		line, err = reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
		currentLineNum++
	}

	fmt.Print(line)
}
