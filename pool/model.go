package pool

import (
	"context"
)

type Pool struct {
	sample []rune
	hash   [16]byte

	maxWorkers int
	blocksize  int
	n          int

	ctx    context.Context
	cancel context.CancelFunc

	lg *logger
}
