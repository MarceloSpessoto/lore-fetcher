package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
)

type Logger struct {
	mu       sync.Mutex
	entries  []string
	onChange func()
	stdout   io.Writer
}

func New() *Logger {
	return &Logger{}
}

func NewWithStdout() *Logger {
	return &Logger{stdout: os.Stdout}
}

func (l *Logger) SetOnChange(fn func()) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.onChange = fn
}

func (l *Logger) Write(p []byte) (n int, err error) {
	line := string(bytes.TrimRight(p, "\n"))
	l.mu.Lock()
	if line != "" {
		l.entries = append(l.entries, line)
	}
	onChange := l.onChange
	stdout := l.stdout
	l.mu.Unlock()
	if stdout != nil && line != "" {
		fmt.Fprintln(stdout, line)
	}
	if onChange != nil {
		onChange()
	}
	return len(p), nil
}

func (l *Logger) Entries() []string {
	l.mu.Lock()
	defer l.mu.Unlock()
	result := make([]string, len(l.entries))
	copy(result, l.entries)
	return result
}
