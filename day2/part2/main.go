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
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 1000005), 1000005)
}
func getNextString(scanner *bufio.Scanner) (string, error) {
	scanned := scanner.Scan()
	if !scanned {
		return "", errors.New("empty")
	}
	return scanner.Text(), nil
}
func getNextInt(scanner *bufio.Scanner) (int, error) {
	s, err := getNextString(scanner)
	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}
func getNextInt64(scanner *bufio.Scanner) (int64, error) {
	s, err := getNextString(scanner)
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}
func getNextFloat64(scanner *bufio.Scanner) (float64, error) {
	s, err := getNextString(scanner)
	if err != nil {
		return 0, err
	}
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
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
	policies := parse(scanner)
	for _, p := range policies {
		fmt.Println(p)
	}
	c := countValidPolicies(policies)
	fmt.Fprintln(writer, c)
}
type policy struct {
	pos1 int
	pos2 int
	character string
	password string
}
func parse(scanner *bufio.Scanner) []policy {
	var policies []policy
	for ;; {
		r, err := getNextString(scanner)
		if err != nil {
			break
		}
		s := strings.Split(r, "-")
		pos1, _ := strconv.Atoi(string(s[0]))
		pos2, _ := strconv.Atoi(string(s[1]))

		c, err := getNextString(scanner)
		if err != nil {
			break
		}
		cRune := []rune(c)
		character := string(cRune[0])

		password, err := getNextString(scanner)
		if err != nil {
			break
		}
		
		p := policy{pos1, pos2, character, password}
		policies = append(policies, p)
	}
	return policies
}
func countValidPolicies(policies []policy) int {
	validCount := 0
	for _, p := range policies {
		passwordRune := []rune(p.password)
		con1 := (string(passwordRune[p.pos1 - 1]) == p.character)
		con2 := (string(passwordRune[p.pos2 - 1]) == p.character)
		if (con1 || con2) && !(con1 && con2) {
			validCount++
		}
	}
	return validCount
}