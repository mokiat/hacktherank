# Hack The Rank

![Build Status](https://github.com/mokiat/hacktherank/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/mokiat/hacktherank)](https://goreportcard.com/report/github.com/mokiat/hacktherank)

A tool that should make it easier to write and test Go solutions for HackerRank challenges.
The `hacktherank` tool allows you to run a set of test data samples against your source code.

## Usage

1. Install the tool

	```sh
	go install github.com/mokiat/hacktherank@latest
	```

1. Create a test data file

	This is a file which contains the input lines to test with and the expected output lines. It is possible to specify multiple test scenarios (or samples) within a single file. Following is an example for a sum challenge.

	```
	input:
	3 5
	output:
	8
	input:
	10 20
	output:
	30

	```

	**Note:** Make sure that all relevant new-line symbols are added and that no unexpected blank space characters are inserted. The tool matches lines exactly - without trimming.

1. Write your solution

	You should create a `main.go` file somwhere and add your Go implementation to it.

1. Test your solution with the `hacktherank` tool

	Following is a demonstration using the example folder of this project.

	```sh
	hacktherank -main example/main.go -data example/data.txt
	```

1. Iterate over your solution

	Once you start tackling more complicated challenges, it is unlikely that you will be able to solve them on the first try.

	Make adjustments to your `main.go` file and just rerun the command to see if you have solved it.
