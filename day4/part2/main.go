package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
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
	// fmt.Println("byr", p.isValidByr())
	// fmt.Println("Iyr", p.isValidIyr())
	// fmt.Println("Eyr", p.isValidEyr())
	// fmt.Println("Hgt", p.isValidHgt())
	// fmt.Println("Ecl", p.isValidEcl())
	// fmt.Println("Pid", p.isValidPid())
	// fmt.Println("Hcl", p.isValidHcl())
	// fmt.Println("Cid", p.isValidCid())
	return p.isValidByr() && p.isValidIyr() && p.isValidEyr() && p.isValidHgt() && p.isValidEcl() && p.isValidPid() && p.isValidHcl() && p.isValidCid()
}

func (p passport) isValidByr() bool {
	target := p.byr
	if target == "" {
		return false
	}
	if len(target) != 4 {
		return false
	}
	i, err := strconv.Atoi(target)
	if err != nil {
		return false
	}
	if i < 1920 || 2002 < i {
		return false
	}
	return true
}

func (p passport) isValidIyr() bool {
	target := p.iyr
	if target == "" {
		return false
	}
	if len(target) != 4 {
		return false
	}
	i, err := strconv.Atoi(target)
	if err != nil {
		return false
	}
	if i < 2010 || 2020 < i {
		return false
	}
	return true
}

func (p passport) isValidEyr() bool {
	target := p.eyr
	if target == "" {
		return false
	}
	if len(target) != 4 {
		return false
	}
	i, err := strconv.Atoi(target)
	if err != nil {
		return false
	}
	if i < 2020 || 2030 < i {
		return false
	}
	return true
}


func (p passport) isValidHgt() bool {
	target := p.hgt
	if target == "" {
		return false
	}

	cmRegex := regexp.MustCompile(`^(\d*)cm$`)
	inRegex := regexp.MustCompile(`^(\d*)in$`)
	if cmRegex.MatchString(target) {
		match := cmRegex.FindStringSubmatch(target)
		if len(match) != 0 {
			hgt, err := strconv.Atoi(match[1])
			if err != nil {
				return false
			}
			if hgt < 150 || 193 < hgt {
				return false
			}
		}
	} else if inRegex.MatchString(target) {
		match := inRegex.FindStringSubmatch(target)
		if len(match) != 0 {
			hgt, err := strconv.Atoi(match[1])
			if err != nil {
				return false
			}
			if hgt < 59 || 76 < hgt {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func (p passport) isValidHcl() bool {
	target := p.hcl
	if target == "" {
		return false
	}
	hclRegex := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	return hclRegex.MatchString(target)
}

func (p passport) isValidEcl() bool {
	target := p.ecl
	if target == "" {
		return false
	}
	eclRegex := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	return eclRegex.MatchString(target)
}

func (p passport) isValidPid() bool {
	target := p.pid
	if target == "" {
		return false
	}
	pidRegex := regexp.MustCompile(`^[0-9]{9}$`)
	return pidRegex.MatchString(target)
}

func (p passport) isValidCid() bool {
	return true
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