package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func configure(scanner *bufio.Scanner) {
	//scanner.Split(bufio.ScanWords)
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
	var nums []int
	for ;; {
		n, err := getNextInt(scanner)
		if err != nil {
			break
		}
		nums = append(nums, n)
	}
	ans := -1
	for x := 0; x < len(nums) - 1; x++ {
		rest := 2020 - nums[x]
		for y := x + 1; y < len(nums); y++ {
			if rest == nums[y] {
				ans = nums[y] * nums[x]
				break
			}
		}
		if ans > 0 {
			break
		}
	}
	fmt.Fprintln(writer, ans)
}
