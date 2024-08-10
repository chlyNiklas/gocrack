package pool

import (
	"fmt"
	"unicode/utf8"
)

type logger struct {
	logChan chan string
}

func (l logger) Close() {
	close(l.logChan)
	fmt.Printf("\n")
}
func newLogger() *logger {
	l := &logger{
		make(chan string, 10),
	}
	go l.logRutine()

	return l
}

func (l *logger) logRutine() {
	for msg := range l.logChan {
		fmt.Printf("\rDigits: %d \t %s", utf8.RuneCountInString(msg), msg)
	}
}

func (l *logger) log(msg string) {
	l.logChan <- msg
}
