package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mokiat/hacktherank/internal/testapp"
	"github.com/mokiat/hacktherank/internal/testdata"
)

var (
	mainFile = flag.String("main", "main.go", "go executable main file")
	dataFile = flag.String("data", "data.txt", "test data file")
)

func main() {
	flag.Parse()

	log.Println("Testing...")
	if err := runTool(); err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Println("All samples passed successfully")
}

func runTool() error {
	samples, err := readSamples(*dataFile)
	if err != nil {
		return fmt.Errorf("failed to read samples: %w", err)
	}

	executableFilename, err := testapp.Build(*mainFile)
	if err != nil {
		return fmt.Errorf("failed to build program: %w", err)
	}

	for sampleIndex, sample := range samples {
		if err := runSample(executableFilename, sample); err != nil {
			return fmt.Errorf("sample %d failure: %w", sampleIndex+1, err)
		}
	}
	return nil
}

func readSamples(filename string) ([]testdata.Sample, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open testdata file: %w", err)
	}
	defer file.Close()

	samples, err := testdata.Parse(file)
	if err != nil {
		return nil, fmt.Errorf("failed to parse testdata: %w", err)
	}
	return samples, nil
}

func runSample(executableFilename string, sample testdata.Sample) error {
	program, err := testapp.Run(executableFilename)
	if err != nil {
		return fmt.Errorf("failed to start program: %w", err)
	}
	defer program.Close()

	for _, inputLine := range sample.InputLines {
		if _, err := fmt.Fprintln(program.StdIn(), inputLine); err != nil {
			return fmt.Errorf("failed to send input line to program: %w", err)
		}
	}

	stdOutScanner := bufio.NewScanner(program.StdOut())
	for _, expectedOutputLine := range sample.OutputLines {
		if !stdOutScanner.Scan() {
			return fmt.Errorf("program did not produce output line: waited on %q but got nothing", expectedOutputLine)
		}
		if err := stdOutScanner.Err(); err != nil {
			return fmt.Errorf("failed to read output line from program: %w", err)
		}
		if outputLine := stdOutScanner.Text(); outputLine != expectedOutputLine {
			return fmt.Errorf("output mismatch: expected %q but got %q", expectedOutputLine, outputLine)
		}
	}
	if stdOutScanner.Scan() {
		return fmt.Errorf("program produced excess output %q", stdOutScanner.Text())
	}
	return nil
}
