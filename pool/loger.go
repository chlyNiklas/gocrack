package pool

import (
	"fmt"
	"unicode/utf8"
)

type logger struct {
	logChan chan string
	enabled bool
}

func (l logger) Close() {
	close(l.logChan)
	if l.enabled {
		fmt.Printf("\n")
	}
}
func newLogger(enabled bool) *logger {
	l := &logger{
		make(chan string, 10),
		enabled,
	}
	if enabled {
		go l.logRutine()
	}

	return l
}

func (l *logger) logRutine() {
	for msg := range l.logChan {
		fmt.Printf("\rDigits: %d \t %s", utf8.RuneCountInString(msg), msg)
	}
}

func (l *logger) log(msg string) {
	if l.enabled {
		l.logChan <- msg

	}
}
