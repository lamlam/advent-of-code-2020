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
	program := []instruction{}
	for {
		s, err := getNextLine(scanner)
		if err != nil {
			break
		}
		ins := parseInstruction(s)
		program = append(program, ins)
	}
	fmt.Fprintln(writer, execCode(program))
}

type instruction struct {
	op string
	arg int
	call int
}

func parseInstruction(line string) instruction {
	splitted := strings.Split(line, " ")
	arg, _ := strconv.Atoi(splitted[1])
	return instruction{splitted[0], arg, 0}
}

func execCode(program []instruction) int {
	value := 0
	for pos := 0; pos < len(program); {
		ins := &program[pos]
		
		if ins.call > 0 {
			break
		}

		ins.call++

		//fmt.Println(pos, ins, value)
		switch ins.op {
		case "nop":
			pos++
		case "jmp":
			pos += ins.arg
		case "acc":
			value += ins.arg
			pos++
		default:
			fmt.Println("Unknown instruction")
			continue
		}
	}
	return value
}
