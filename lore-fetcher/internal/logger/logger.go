package logger

import (
	"bytes"
	"sync"
)

type Logger struct {
	mu       sync.Mutex
	entries  []string
	onChange func()
}

func New() *Logger {
	return &Logger{}
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
	l.mu.Unlock()
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
