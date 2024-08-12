package pool

import (
	"context"
	"log"

	"github.com/chlyniklas/gocrack/solver"
)

type job struct {
	from int
	to   int
}

func (p *Pool) Crack() string {
	jobs := make(chan job)
	result := make(chan string)

	p.lg = newLogger(p.logging)

	p.ctx, p.cancel = context.WithCancel(context.Background())

	go p.employer(jobs)
	for range p.maxWorkers - 1 {
		go p.worker(jobs, result)
	}
	go p.loggingWorker(jobs, result)

	if p.logging {
		log.Printf("Created %d workers with block size: %d", p.maxWorkers, p.blocksize)
	}

	res := <-result

	p.cancel()
	p.lg.Close()
	return res
}

func (p *Pool) worker(jobs <-chan job, result chan<- string) {
	s := solver.New(p.sample, p.checkFunction)
	for {
		select {
		case j := <-jobs:
			// test the block given by the job
			for i := j.from; i < j.to; i++ {
				password, ok := s.CheckStringAtPosition(i)
				if ok {
					result <- password
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

	s := solver.New(p.sample, p.checkFunction)

	for {
		select {
		case j := <-jobs:
			p.lg.log(s.CreateUniqueString(j.from))
			for i := j.from; i < j.to; i++ {
				password, ok := s.CheckStringAtPosition(i)
				// log.Println(password)
				if ok {
					result <- password
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
			p.n = from
			jobs <- job{from, from + p.blocksize}
		}
	}
}
