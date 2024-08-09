package main

import (
	"crypto/md5"
	"log"
	"testing"
	"time"

	"github.com/chlyniklas/gocrack/pool"
	"github.com/chlyniklas/gocrack/solver"
)

var hash [16]byte = md5.Sum([]byte("psswrd"))
var sample []rune = []rune("abcdefghijklmnopqrstuvwxyz")

func Benchmark_NoBlocks_4Workers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jobs := make(chan int, 10)
		result := make(chan string)

		s := solver.New(sample, hash)

		go func() {
			for i := 0; true; i++ {
				jobs <- i
			}
		}()

		worker := func() {
			for j := range jobs {
				password, ok := s.CheckStringAtPosition(j)
				// log.Println(password)
				if ok {
					result <- password
				}
			}
		}

		go worker()
		go worker()
		go worker()
		go worker()

		<-result

	}
	log.Println()
	log.Println("number of iterations: ", b.N)
	log.Println("elapsed:", b.Elapsed()/time.Duration(b.N))
}
func Benchmark_Poo_Blocksize500_4Workers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := pool.New(sample, hash)
		_ = p.ConfigureWorkers(4, 500)
		p.Crack()
	}

	log.Println()
	log.Println("number of iterations: ", b.N)
	log.Println("elapsed:", b.Elapsed()/time.Duration(b.N))
}
