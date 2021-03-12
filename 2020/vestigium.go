package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type answer struct {
	trace int
	rRow  int
	rCol  int
}

func main() {
	var nrTestCases int
	var sizeOfMatrices []int
	var matrices [][][]int

	readInput := bufio.NewScanner(os.Stdin)

	readInput.Scan()
	nrTestCases, err := strconv.Atoi(readInput.Text())
	if err != nil {
		log.Println(err)
	}

	sizeOfMatrices = make([]int, nrTestCases)
	matrices = make([][][]int, nrTestCases)

	for i := 0; i < nrTestCases; i++ {
		readInput.Scan()

		sizeOfMatrices[i], err = strconv.Atoi(readInput.Text())
		matrices[i] = make([][]int, sizeOfMatrices[i])
		if err != nil {
			log.Println(err)
		}

		currentAns := answer{}

		colChecks := make([]map[int]interface{}, sizeOfMatrices[i])
		for k := range colChecks {
			colChecks[k] = make(map[int]interface{}, sizeOfMatrices[i])
		}
		colDuplicated := make([]bool, sizeOfMatrices[i])

		for j := 0; j < sizeOfMatrices[i]; j++ {
			readInput.Scan()

			if err != nil {
				log.Println(err)
			}

			rowList := strings.Split(readInput.Text(), " ")
			matrices[i][j] = make([]int, sizeOfMatrices[i])

			rowCheck := make(map[int]interface{}, sizeOfMatrices[i])
			rowDuplicated := false
			for k, r := range rowList {
				rowValue, err := strconv.Atoi(r)
				if err != nil {
					log.Println(err)
				}

				if _, ok := rowCheck[rowValue]; ok && !rowDuplicated {
					currentAns.rRow++
					rowDuplicated = true
				}
				rowCheck[rowValue] = struct{}{}

				if k == j {
					currentAns.trace += rowValue
				}

				if _, ok := colChecks[k][rowValue]; ok && !colDuplicated[k] {
					colDuplicated[k] = true
					currentAns.rCol++
				}
				colChecks[k][rowValue] = struct{}{}

				matrices[i][j][k] = rowValue
			}
		}

		output(i, currentAns.trace, currentAns.rRow, currentAns.rCol)
	}
}

func output(t, n, r, c int) {
	fmt.Printf("Case #%d: %d %d %d\n", t+1, n, r, c)
}
