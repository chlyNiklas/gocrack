package pool

import (
	"context"

	"github.com/chlyniklas/gocrack/solver"
)

type Pool struct {
	sample        []rune
	checkFunction solver.CheckFunction

	maxWorkers int
	blocksize  int
	n          int

	ctx    context.Context
	cancel context.CancelFunc

	lg      *logger
	logging bool
}
