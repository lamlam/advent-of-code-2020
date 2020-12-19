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
	forms := readFormOfAllGoroups(scanner)
	total := 0
	for _, form := range forms {
		current := 0
		for _, yes := range form {
			if yes {
				total++
				current++
			}
		}
		fmt.Println(form, current)
	}
	fmt.Println(total)
}

type form []bool

func readFormOfAllGoroups(scanner *bufio.Scanner) []form {
	formStrings := []string{}
	forms := []form{}
	for {
		line, err := getNextLine(scanner)
		if err != nil {
			forms = append(forms, createFormOfGroup(formStrings))
			break
		}

		if line == "" {
			forms = append(forms, createFormOfGroup(formStrings))
			formStrings = []string{}
		} else {
			formStrings = append(formStrings, line)
		}
	}
	return forms
}

func createFormOfGroup(formStrings []string) form {
	yesCounts := make([]int, 26)
	for _, formString := range formStrings {
		for i := range formString {
			yesCounts[formString[i] - "a"[0]]++
		}
	}
	f := make(form, 26)
	for i, count := range yesCounts {
		if count == len(formStrings) {
			f[i] = true
		}
	}
	return f
}