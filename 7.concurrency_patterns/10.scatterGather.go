package concurrencypatterns

import (
	"fmt"
	"io"
	"strings"
)

// https://medium.com/star-gazers/go-scatter-gather-pattern-a439c70afe16

func ScatterGather(concurrency int, jobs []int) {
	errors := make(chan error)
	results := make(chan string)
	// Scatter
	for i := 0; i < len(jobs); i += concurrency {
		go func(i int) {
			end := i + concurrency
			if end > len(jobs) {
				end = len(jobs)
			}
			res, err := process(jobs[i:end])
			if err != nil {
				errors <- err
				return
			}
			results <- res
		}(i)
	}
	// Gather
	for i := 0; i < len(jobs); i += concurrency {
		select {
		case err := <-errors:
			fmt.Println("Error", err)
		case result := <-results:
			fmt.Println("Success", result)
		}
	}
}

// just some trash to make strings and errors
func process(jobs []int) (string, error) {
	if len(jobs) == 0 {
		return "", io.EOF
	}
	if jobs[0]%2 == 0 {
		return "", fmt.Errorf("fake Error")
	}
	var sb strings.Builder
	for _, n := range jobs {
		sb.WriteString(fmt.Sprintf("%d", n))
	}
	return sb.String(), nil
}
