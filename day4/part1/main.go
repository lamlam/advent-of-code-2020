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
func getNextInt(scanner *bufio.Scanner) (int, error) {
	s, err := getNextLine(scanner)
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
	s, err := getNextLine(scanner)
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
	s, err := getNextLine(scanner)
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
	passports := readPassports(scanner)
	validCount := 0
	for _, p := range passports {
		if p.isValid() {
			validCount++
		}
	}
	fmt.Fprintln(writer, validCount)
}
type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	ecl string
	pid string
	cid string
	hcl string
}

func (p passport) isValid() bool {
	return (p.byr != "") && (p.iyr != "") && (p.eyr != "") && (p.hgt != "") && (p.ecl != "") && (p.pid != "") && (p.hcl != "")
}

func readPassports(scanner *bufio.Scanner) []passport {
	passportString := ""
	passports := []passport{}
	for s, err := getNextLine(scanner); err == nil; s, err = getNextLine(scanner) {
		if s != "" {	
			passportString = passportString + s + " "
			continue
		}
		p := initPassport(passportString)
		passports = append(passports, p)
		passportString = ""
	}
	return passports
}

func initPassport(passportString string) passport {
	p := passport{}
	passportString = strings.TrimRight(passportString, " ")	
	data := strings.Split(passportString, " ")
	for _, s := range data {
		kv := strings.Split(s, ":")
		k := kv[0]
		v := kv[1]
		switch k {
		case "byr":
			p.byr = v
		case "iyr":
			p.iyr = v
		case "eyr":
			p.eyr = v
		case "hgt":
			p.hgt = v
		case "ecl":
			p.ecl = v
		case "pid":
			p.pid = v
		case "cid":
			p.cid = v
		case "hcl":
			p.hcl = v
		default:
			fmt.Println("unknown key", k)
		}
	}
	return p
}