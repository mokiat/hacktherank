package testapp

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Run(filename string) (*Executable, error) {
	cmd := exec.Command(filename)

	stdIn, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to establish stdin pipe: %w", err)
	}
	stdOut, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to establish stdout pipe: %w", err)
	}
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start executable: %w", err)
	}

	return &Executable{
		cmd:    cmd,
		stdIn:  stdIn,
		stdOut: stdOut,
	}, nil
}

type Executable struct {
	cmd    *exec.Cmd
	stdIn  io.WriteCloser
	stdOut io.ReadCloser
}

func (e *Executable) StdIn() io.Writer {
	return e.stdIn
}

func (e *Executable) StdOut() io.Reader {
	return e.stdOut
}

func (e *Executable) Close() error {
	if err := e.cmd.Process.Kill(); err != nil {
		return fmt.Errorf("failed to kill executable: %w", err)
	}
	return nil
}
