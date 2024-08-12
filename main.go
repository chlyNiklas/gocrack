package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/chlyniklas/gocrack/pool"
)

type args struct {
	Hash       [16]byte
	Sample     []rune
	MaxWorkers int
	Blocksize  int
}

func main() {
	a, err := GetArgs()

	if err != nil {
		fmt.Println(err)
	}

	workerPool := pool.New(a.Sample, a.Hash)
	workerPool.SetMaxWorkers(a.MaxWorkers)
	workerPool.SetBlocksize(a.Blocksize)

	log.Println("Cracking hash: ", hex.EncodeToString(a.Hash[:]))
	tStart := time.Now()
	log.Println("With chars:", string(a.Sample))

	fmt.Printf("%s: %s \n", hex.EncodeToString(a.Hash[:]), workerPool.Crack())
	timeTaken := time.Now().Sub(tStart)
	fmt.Printf("Cracked in %s\n", timeTaken)
	// fmt.Printf("With %.0f hashes/second\n", float64(workerPool.GetNumberOfHashesChecked())/timeTaken.Seconds())
	fmt.Printf("With %.5v hashes/second\n", float64(workerPool.GetNumberOfHashesChecked())/timeTaken.Seconds())
}

func GetArgs() (a args, err error) {

	flag.IntVar(&a.MaxWorkers, "workers", 0, "")
	flag.IntVar(&a.Blocksize, "blocksize", 0, "")
	sampleString := flag.String("sample", "abcdefghijklmnopqrstuvwxyz", "")

	flag.Parse()

	log.Println(a.MaxWorkers)

	a.Sample = []rune(*sampleString)

	// Read hash as argument
	a.Hash = md5.Sum([]byte("niklas"))
	if hashString := flag.Arg(0); hashString != "" {
		val, err := hex.DecodeString(hashString)
		if err != nil {
			return a, errors.New("Hash must be in hex")
		}
		a.Hash = [16]byte(val)
	}
	return
}
