package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	impossible = "IMPOSSIBLE"
	cameron    = "C"
	jamie      = "J"
)

type StartEnd struct {
	start int
	end   int
	str   *string
}

type Parent struct {
	name     string
	freeTime int
}

func main() {
	var nrTestCases int
	var nrActivities []int
	var startEnd [][]StartEnd
	var startEndOri [][]StartEnd

	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()

	nrTestCases, _ = strconv.Atoi(reader.Text())
	nrActivities = make([]int, nrTestCases)
	startEnd = make([][]StartEnd, nrTestCases)
	startEndOri = make([][]StartEnd, nrTestCases)

	parentJ := Parent{
		name: jamie,
	}
	parentC := Parent{
		name: cameron,
	}

	for i := 0; i < nrTestCases; i++ {
		reader.Scan()
		nrActivities[i], _ = strconv.Atoi(reader.Text())
		startEnd[i] = make([]StartEnd, nrActivities[i])
		startEndOri[i] = make([]StartEnd, nrActivities[i])

		for j := 0; j < nrActivities[i]; j++ {
			reader.Scan()
			startEndSplit := strings.Split(reader.Text(), " ")
			startEnd[i][j].start, _ = strconv.Atoi(startEndSplit[0])
			startEnd[i][j].end, _ = strconv.Atoi(startEndSplit[1])
			startEnd[i][j].str = new(string)
			startEndOri[i][j].str = startEnd[i][j].str
		}

		sort.Slice(startEnd[i], func(k, j int) bool {
			return startEnd[i][k].start < startEnd[i][j].start
		})

		for j := 0; j < nrActivities[i]; j++ {
			if parentJ.freeTime <= startEnd[i][j].start {
				*startEnd[i][j].str += parentJ.name
				parentJ.freeTime = startEnd[i][j].end
			} else if parentC.freeTime <= startEnd[i][j].start {
				*startEnd[i][j].str += parentC.name
				parentC.freeTime = startEnd[i][j].end
			} else {
				*startEnd[i][j].str = impossible
				break
			}
		}

		parentJ.freeTime = 0
		parentC.freeTime = 0
	}

	for i := 0; i < nrTestCases; i++ {
		str := ""
		for _, v := range startEndOri[i] {
			if *v.str == impossible {
				str = impossible
				break
			}
			str += *v.str
		}
		outputWrite(i+1, str)
	}
}

func outputWrite(i int, str string) {
	fmt.Printf("Case #%d: %s\n", i, str)
}
