package pool

import (
	"context"

	"github.com/chlyniklas/gocrack/checkfunctionfactory"
)

type Pool struct {
	charSet       []byte
	checkFunction checkfunctionfactory.CheckFunction

	maxWorkers int
	blocksize  int
	n          int

	ctx    context.Context
	cancel context.CancelFunc

	lg      *logger
	logging bool
}
