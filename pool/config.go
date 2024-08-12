package pool

import (
	"errors"
	"runtime"

	"github.com/chlyniklas/gocrack/solver"
)

func New(sample []rune, checkFunction solver.CheckFunction) *Pool {
	return &Pool{
		sample:        sample,
		maxWorkers:    runtime.NumCPU() * 3,
		blocksize:     50000,
		logging:       true,
		checkFunction: checkFunction,
	}
}

func (p *Pool) ConfigureWorkers(maxWorkers, blocksize int) (err error) {
	if maxWorkers <= 0 {
		return errors.New("maxWorkers must be grater than 0")
	}
	if blocksize <= 0 {
		return errors.New("blocksize must be grater than 0")
	}

	p.blocksize = blocksize
	p.maxWorkers = maxWorkers

	return nil
}

func (p *Pool) SetSample(sample []rune) {
	p.sample = sample
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
