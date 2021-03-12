package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var nrTestCases int
	var inputString []string
	var outputString []string

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	nrTestCases, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Println(err)
	}

	inputString = make([]string, nrTestCases)
	outputString = make([]string, nrTestCases)

	for k := 0; k < nrTestCases; k++ {
		scanner.Scan()
		inputString[k] = scanner.Text()
		openBracketCounter := 0
		var strBuilder strings.Builder

		for _, v := range inputString[k] {
			vint, _ := strconv.Atoi(string(v))

			switch {
			case openBracketCounter < vint:
				openCount := vint - openBracketCounter
				for i := 0; i < openCount; i++ {
					openBracketCounter++
					strBuilder.WriteString("(")
				}
				fallthrough
			case openBracketCounter > vint:
				closeCount := openBracketCounter - vint
				for i := 0; i < closeCount; i++ {
					strBuilder.WriteString(")")
					openBracketCounter--
				}
				fallthrough
			default:
				strBuilder.WriteRune(v)
			}
		}

		for i := openBracketCounter; i != 0; i-- {
			strBuilder.WriteString(")")
		}

		outputString[k] = strBuilder.String()
	}

	for k := 0; k < nrTestCases; k++ {
		output(k+1, outputString[k])
	}
}

func output(i int, str string) {
	fmt.Printf("Case #%d: %s\n", i, str)
}
