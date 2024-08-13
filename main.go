package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/chlyniklas/gocrack/checkfunctionfactory"
	"github.com/chlyniklas/gocrack/pool"
)

type args struct {
	Hash       string
	CheckFunc  checkfunctionfactory.CheckFunction
	charSet    []byte
	MaxWorkers int
	Blocksize  int
}

func main() {
	a, err := getArgs()

	if err != nil {
		fmt.Println(err)
		return
	}

	workerPool := pool.New(a.charSet, a.CheckFunc)
	workerPool.SetMaxWorkers(a.MaxWorkers)
	workerPool.SetBlocksize(a.Blocksize)

	log.Println("Cracking hash: ", a.Hash[:])
	tStart := time.Now()
	log.Println("With chars:", string(a.charSet))

	fmt.Printf("%s: %s \n", a.Hash[:], workerPool.Crack())
	timeTaken := time.Now().Sub(tStart)
	fmt.Printf("Cracked with %.5v hashes/second\n", float64(workerPool.GetNumberOfHashesChecked())/timeTaken.Seconds())
}

func getArgs() (a args, err error) {
	flag.IntVar(&a.MaxWorkers, "workers", 0, "Set the number of go routines that crack")
	flag.IntVar(&a.Blocksize, "blocksize", 0, "Set the number of hashes one worker checks before waiting for a new job")
	charSetString := flag.String("charset", "abcdefghijklmnopqrstuvwxyz", "All characters your password could possibly contain")
	hashType := flag.String("hashType", "md5", "The hash algorithm used to make your hash.\nYou can choose between: md5, sha1, sha256 & sha512")

	flag.Parse()

	a.charSet = []byte(*charSetString)

	a.Hash = flag.Arg(0)
	if a.Hash == "" {
		return a, errors.New("No hash given.")
	}

	cffactory, err := checkfunctionfactory.New(checkfunctionfactory.HashAlgorithm(*hashType), a.Hash)
	if err != nil {
		return a, err
	}

	a.CheckFunc, err = cffactory.CreateCheckFunction()

	return
}
