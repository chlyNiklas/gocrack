package pool

import "context"

type Pool struct {
	sample []rune
	hash   [16]byte

	maxWorkers int
	blocksize  int

	ctx    context.Context
	cancel context.CancelFunc
}
