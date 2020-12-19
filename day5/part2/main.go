package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func configure(scanner *bufio.Scanner) {
	scanner.Split(bufio.ScanLines)
	scanner.Buffer(make([]byte, 1000005), 1000005)
}
func getNextLine(scanner *bufio.Scanner) (string, error) {
	scanned := scanner.Scan()
	if !scanned {
		return "", errors.New("empty")
	}
	return scanner.Text(), nil
}

func main() {
	fp := os.Stdin
	wfp := os.Stdout
	
	scanner := bufio.NewScanner(fp)
	configure(scanner)
	writer := bufio.NewWriter(wfp)
	defer func() {
		r := recover()
		if r != nil {
			fmt.Fprintln(writer, r)
		}
		writer.Flush()
	}()
	solve(scanner, writer)
}

func solve(scanner *bufio.Scanner, writer *bufio.Writer) {
	IDs := make([]bool, 8 * 127 + 8)
	for {
		s, err := getNextLine(scanner)
		if err != nil {
			break
		}
		id := searchBordingPass(s)
		IDs[id] = true
	}
	for i, id := range IDs {
		if !id {
			if i != 0 && i != 128 {
				 if IDs[i-1] && IDs[i+1] {
					fmt.Fprintln(writer, i)
				 }
			}
		}
	}
}

func searchBordingPass(command string) int {
	commandRune := []rune(command)
	first := 0
	last := 127
	for _, c := range commandRune[:7] {
		mid := first + (last - first) / 2
		if string(c) == "F" {
			last = mid
		} else {
			first = mid + 1
		}
		//fmt.Println(first, last, mid)
	}
	row := first

	first = 0
	last = 7
	for _, c := range commandRune[7:] {
		mid := first + (last - first) / 2
		if string(c) == "L" {
			last = mid
		} else {
			first = mid + 1
		}
		//fmt.Println(first, last, mid)
	}
	column := first
	return row * 8 + column
}