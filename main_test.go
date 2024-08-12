package main

import (
	"crypto/md5"
	"log"
	"testing"
	"time"

	"github.com/chlyniklas/gocrack/pool"
)

var hash [16]byte = md5.Sum([]byte("psswrd"))
var sample []rune = []rune("abcdefghijklmnopqrstuvwxyz")

func Benchmark_CrackDefault(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := pool.New(sample, hash)
		p.Crack()
	}

	log.Println("Crack psswrd with default conf")
	log.Println("number of iterations: ", b.N)
	log.Println("elapsed:", b.Elapsed()/time.Duration(b.N))
}
