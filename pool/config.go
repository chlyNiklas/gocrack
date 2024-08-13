package pool

import (
	"errors"
	"runtime"

	"github.com/chlyniklas/gocrack/checkfunctionfactory"
)

func New(charSet []byte, checkFunction checkfunctionfactory.CheckFunction) *Pool {
	return &Pool{
		charSet:       charSet,
		maxWorkers:    runtime.NumCPU() * 3,
		blocksize:     50000,
		logging:       true,
		checkFunction: checkFunction,
	}
}

// Set's the rune set
func (p *Pool) SetRuneSet(charSet []byte) {
	p.charSet = charSet
}
func (p *Pool) SetMaxWorkers(maxWorkers int) (err error) {
	if maxWorkers <= 0 {
		return errors.New("Workers must be grater than 0")
	}

	p.maxWorkers = maxWorkers

	return nil

}
func (p *Pool) SetLogging(enabled bool) {
	p.logging = enabled
}

func (p *Pool) SetBlocksize(blocksize int) (err error) {
	if blocksize <= 0 {
		return errors.New("blocksize must be grater than 0")
	}

	p.blocksize = blocksize

	return nil

}

func (p *Pool) GetNumberOfHashesChecked() int {
	return p.n * p.blocksize
}
