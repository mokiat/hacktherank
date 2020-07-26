# Hack The Rank

A tool that should make it easier to write and test Go solutions for HackerRank challenges.
The `hacktherank` tool allows you to run a set of test data samples against your source code.

## Usage

1. First you need to download the tool

	```sh
	GO111MODULE=on go get github.com/mokiat/hacktherank
	```

	This should install the `hacktherank` executable in your `$GOPATH/bin` directory. Make sure you have
	that directory added to your `PATH` environment variable.

1. Next, you need to create a test data file 

	This is a file which contains the input lines to test with and the expected output lines. 
	It is possible to specify multiple test scenarios (or samples) within a single file. Following is an 
	example for a sum challenge.

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

	**Note:** Make sure that all relevant new-line symbols are added and that no unexpected blank space characters are inserted. The tool matches lines exactly, without trimming.

1. You need to write your solution

	You should create a `main.go` file somwhere and add your Go implementation to it.

1. You need to run the `hacktherank` tool.

	Following is a demonstration using the example folder of this project.

	```sh
	hacktherank -main example/main.go -data example/data.txt
	```

1. Iterate over your solution

	Once you start tackling more complicated challenges, it is unlikely that you will be able to solve them on the first run.

	Make adjustments to your `main.go` file and just rerun the command to see if you have solved it.


## Snippets

Following are snippets of common code fragments I have had to use for my solutions.

```go
func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
```

```go
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```

```go
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
