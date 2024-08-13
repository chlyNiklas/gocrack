package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/chlyniklas/gocrack/pool"
	"github.com/chlyniklas/gocrack/solver"
)

type args struct {
	Hash       string
	CheckFunc  solver.CheckFunction
	charSet    []byte
	MaxWorkers int
	Blocksize  int
}

func main() {
	a, err := GetArgs()

	if err != nil {
		fmt.Println(err)
	}

	workerPool := pool.New(a.charSet, a.CheckFunc)
	workerPool.SetMaxWorkers(a.MaxWorkers)
	workerPool.SetBlocksize(a.Blocksize)

	log.Println("Cracking hash: ", a.Hash[:])
	tStart := time.Now()
	log.Println("With chars:", string(a.charSet))

	fmt.Printf("%s: %s \n", a.Hash[:], workerPool.Crack())
	timeTaken := time.Now().Sub(tStart)
	fmt.Printf("Cracked in %s\n", timeTaken)
	// fmt.Printf("With %.0f hashes/second\n", float64(workerPool.GetNumberOfHashesChecked())/timeTaken.Seconds())
	fmt.Printf("With %.5v hashes/second\n", float64(workerPool.GetNumberOfHashesChecked())/timeTaken.Seconds())
}

func GetArgs() (a args, err error) {

	flag.IntVar(&a.MaxWorkers, "workers", 0, "Set the number of go routines that crack")
	flag.IntVar(&a.Blocksize, "blocksize", 0, "Set the number of hashes one worker checks before waiting for a new job")
	charSetString := flag.String("charset", "abcdefghijklmnopqrstuvwxyz", "All characters your password could possibly contain")
	hashType := flag.String("hashType", "md5", "The hash algorithm used to make your hash. You can choose between md5 & sha256")

	flag.Parse()

	log.Println(a.MaxWorkers)

	a.charSet = []byte(*charSetString)

	// Read hash as argument
	var hashSlice []byte
	a.Hash = flag.Arg(0)
	if a.Hash != "" {
		val, err := hex.DecodeString(a.Hash)
		if err != nil {
			return a, errors.New("Hash must be in hex")
		}
		hashSlice = val
	}
	// set hashtype
	switch *hashType {
	case "md5":
		hash := [16]byte(hashSlice)
		a.CheckFunc = func(b []byte) bool {
			return md5.Sum(b) == hash
		}
	case "sha256":
		hash := [32]byte(hashSlice)
		a.CheckFunc = func(b []byte) bool {
			return sha256.Sum256(b) == hash
		}

	}
	return
}
