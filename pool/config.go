package pool

import (
	"errors"
	"runtime"
)

func New(sample []rune, hash [16]byte) *Pool {
	return &Pool{
		sample:     sample,
		hash:       hash,
		maxWorkers: runtime.NumCPU() * 3,
		blocksize:  runtime.NumCPU() * 100,
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

func (p *Pool) SetHash(hash [16]byte) {
	p.hash = hash
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
