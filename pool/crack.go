package pool

import (
	"context"
	"log"
	"unicode/utf8"

	"github.com/chlyniklas/gocrack/solver"
)

type job struct {
	from int
	to   int
}

func (p *Pool) Crack() string {
	jobs := make(chan job, 10)
	result := make(chan string)

	p.ctx, p.cancel = context.WithCancel(context.Background())

	go p.employer(jobs)
	for range p.maxWorkers - 1 {
		go p.worker(jobs, result)
	}
	go p.loggingWorker(jobs, result)

	log.Printf("Created %d workers with block size: %d", p.maxWorkers, p.blocksize)

	return <-result
}

func (p *Pool) worker(jobs <-chan job, result chan<- string) {
	s := solver.New(p.sample, p.hash)
	for {
		select {
		case j := <-jobs:
			// test the block given by the job
			for i := j.from; i < j.to; i++ {
				password, ok := s.CheckStringAtPosition(i)
				// log.Println(password)
				if ok {
					result <- password
					p.cancel()
					return
				}
			}
		case <-p.ctx.Done():
			// if the password is found: stop worker
			return
		}

	}
}

func (p *Pool) loggingWorker(jobs <-chan job, result chan<- string) {

	s := solver.New(p.sample, p.hash)

	for {
		select {
		case j := <-jobs:
			str := s.CreateUniqueString(j.from)
			log.Printf("Digits: %d \t %s", utf8.RuneCountInString(str), str)
			for i := j.from; i < j.to; i++ {
				password, ok := s.CheckStringAtPosition(i)
				// log.Println(password)
				if ok {
					result <- password
					p.cancel()
					return
				}
			}
		case <-p.ctx.Done():
			return
		}
	}

}

func (p *Pool) employer(jobs chan<- job) {
	for i := 0; true; i++ {
		select {
		case <-p.ctx.Done():
			return
		default:
			from := i * p.blocksize
			jobs <- job{from, from + p.blocksize}
		}
	}
}
