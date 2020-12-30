package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	bags := map[string]bag{}
	for {
		s, err := getNextLine(scanner)
		if err != nil {
			break
		}
		b := parseBag(s)
		bags[b.color] = b
	}
	fmt.Fprintln(writer, countBag(bags["shiny gold"], bags))
}

type bag struct {
	color string
	contain map[string]int
	numBags int
	isSearched bool
}

func parseBag(line string) bag {
	replaced := strings.ReplaceAll(line, ",", "")
	replaced = strings.ReplaceAll(replaced, ".", "")
	splited := strings.Split(replaced, " ")

	b := bag{}
	b.color = strings.Join(splited[:2], " ")

	if strings.Contains(line, "no other bags") {
		return b
	}

	b.contain = map[string]int{}
	containStrings := splited[4:]

	for {
		num, _ := strconv.Atoi(containStrings[0])
		b.contain[strings.Join(containStrings[1:3], " ")] = num
		
		if len(containStrings) <= 4 {
			break
		}
		containStrings = containStrings[4:]
	}
	return b
}

func countBag(b bag, bags map[string]bag) int {
	count := 0
	if b.isSearched {
		return b.numBags
	}
	if len(b.contain) == 0 {
		b.isSearched = true
		b.numBags = 0
		return b.numBags
	}
	for k, v := range b.contain {
		count += countBag(bags[k], bags) * v + v
	}
	b.isSearched = true
	b.numBags = count
	return b.numBags
}
