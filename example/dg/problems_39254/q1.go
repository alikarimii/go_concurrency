package problems_39254

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
	"sync"
)

// github.com/alikarimii/go_concurrency/example/dg/problems_39254
// test file exist in this repo
func Run(parser InputParser, input []byte) bool {
	_, storeCoverage, start, end := parser.Parse(input)
	result := true
	for v := range check(start, end, storeCoverage) {
		result = result && v
	}

	return result
}

func NewParser() InputParser {
	return &myParser{}
}
func NewMockParser(in bool) InputParser {
	return &mockParser{ForTest: in}
}

type InputParser interface {
	// input array of byte
	// returns (store_number, array_of_store_coverage, start_district,end_district)
	Parse(b []byte) (int8, [][]int8, int8, int8)
}
type myParser struct{}

func (m *myParser) Parse(b []byte) (int8, [][]int8, int8, int8) {
	reader := bytes.NewReader(b)
	scanner := bufio.NewScanner(reader)
	data := make([]string, 0)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	storeCountStr := strings.Trim(data[0], " ")
	data = data[1:]
	storeCount, _ := strconv.Atoi(storeCountStr)
	coverage := make([][]int8, 0)
	for i := 0; i < storeCount; i++ {
		line := strings.Split(strings.Trim((strings.Trim(data[i], " ")), "\t"), " ")
		start, _ := strconv.Atoi(line[0])
		end, _ := strconv.Atoi(line[1])
		coverage = append(coverage, []int8{int8(start), int8(end)})
	}

	data = data[storeCount:]
	start_district, _ := strconv.Atoi(strings.Trim(data[0], "\t"))
	end_district, _ := strconv.Atoi(strings.Trim(data[1], "\t"))
	return int8(storeCount), coverage, int8(start_district), int8(end_district)
}

type mockParser struct {
	ForTest bool
}

func (m *mockParser) Parse(b []byte) (int8, [][]int8, int8, int8) {
	// true example
	if m.ForTest {
		return 3, [][]int8{{1, 2}, {3, 4}, {5, 6}}, 2, 5
	}
	// false example
	return 2, [][]int8{{1, 10}, {10, 20}}, 21, 21
}

func check(start_district, end_district int8, coverage [][]int8) chan bool {
	out := make(chan bool)
	var wg sync.WaitGroup
	// check district covered by atleast one shop
	checkCoverage := func(district int8) {
		defer wg.Done()
		res := false
		for _, v := range coverage {
			if district >= v[0] && district <= v[1] {
				res = true
				break
			}
		}
		out <- res
	}
	// loop over districts
	for i := start_district; i <= end_district; i++ {
		wg.Add(1)
		go checkCoverage(i)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
