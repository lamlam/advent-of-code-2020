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
	s, _ := createStage(scanner)
	//s.show(writer)
	x, y := 0, 0
	count := 0
	ans := 1
	type slope struct{
		x int
		y int
	}
	for _, sl := range []slope{{1,1},{3,1},{5,1},{7,1},{1,2}} {
		for {
			x, y = move(x, y, sl.x, sl.y)
			res, err := check(s, x, y)
			if err != nil {
				break
			}
			if res {
				count++
			}
		}
		ans = ans * count
		count = 0
		x = 0
		y = 0
	}
	fmt.Fprintln(writer, ans)
}

func createStage(scanner *bufio.Scanner) (stage, error) {
	configure(scanner)
	var s stage
	for {
		l, err := getNextLine(scanner)
		if err != nil {
			break
		}
		var line []bool
		for _, c := range l {
			if string(c) == "#" {
				line = append(line, true)
			} else {
				line = append(line, false)
			}
		}
		s.lines = append(s.lines, line)
	}
	s.width = len(s.lines[0])
	s.height = len(s.lines)
	return s, nil
}

// stage is created from input
// if point of map had tree, point of stage is set to true 
type stage struct {
	lines [][]bool
	width int
	height int
}

func (s stage) show(writer *bufio.Writer) {
	for _, line := range s.lines {
		for _, p := range line {
			if p {
				fmt.Fprintf(writer, "#")
			} else {
				fmt.Fprintf(writer, ".")
			}
		}
		fmt.Fprintf(writer, "\n")
	}
}

func move(x int, y int, xOffset int, yOffset int) (int, int) {
	return x + xOffset, y + yOffset
}

func check(s stage, x int, y int) (bool, error) {
	if s.height <= y {
		return false, errors.New("Reach bottom of stage")
	}
	xx := x % s.width
	//fmt.Println(xx, y, s.lines[y][xx])
	return s.lines[y][xx], nil
}