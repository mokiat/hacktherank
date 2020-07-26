package testdata

import (
	"bufio"
	"fmt"
	"io"
)

func Parse(in io.Reader) ([]Sample, error) {
	scanner := bufio.NewScanner(in)

	var samples []Sample
	var currentSample Sample
	appendSample := func() {
		if currentSample.IsEmpty() {
			return
		}
		samples = append(samples, currentSample)
		currentSample = Sample{}
	}
	addInput := func(line string) {
		currentSample.InputLines = append(currentSample.InputLines, line)
	}
	addOutput := func(line string) {
		currentSample.OutputLines = append(currentSample.OutputLines, line)
	}
	addOperation := addInput

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("failed to scan line: %w", err)
		}
		switch line := scanner.Text(); line {
		case "input:":
			addOperation = addInput
			appendSample()
		case "output:":
			addOperation = addOutput
		default:
			addOperation(line)
		}
	}
	appendSample()

	return samples, nil
}
