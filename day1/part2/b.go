package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
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
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	for i:=0; i<len(nums)-2 && ans < 0; i++ {
		for j:=i; j<len(nums)-1 && ans < 0; j++ {
			for k:=j; k<len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				if 2020 == sum {
					ans = nums[i] * nums[j] * nums[k]
					break
				}
				if sum > 2020 {
					break
				}
			}
		}
	}
	fmt.Fprintln(writer, ans)
}
