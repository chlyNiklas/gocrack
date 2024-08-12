package solver

import (
	"crypto/md5"
)

type Solver struct {
	sample []rune
	hash   [16]byte
}

func New(sample []rune, hash [16]byte) *Solver {
	return &Solver{
		sample,
		hash,
	}
}

func (s *Solver) CheckStringAtPosition(pos int) (str string, ok bool) {
	str = s.CreateUniqueString(pos)
	if s.hash == md5.Sum([]byte(str)) {
		return str, true
	}
	return str, false
}

func (s *Solver) CreateUniqueString(pos int) string {
	nDigits := s.digits(pos)

	pwd := make([]rune, 0, nDigits)
	for range nDigits {
		pwd = append(pwd, s.sample[pos%len(s.sample)])
		pos = pos / len(s.sample)
	}
	return string(pwd)
}

func (s *Solver) digits(pos int) int {
	var sum int
	lastPow := 1
	for i := 1; true; i++ {
		lastPow *= len(s.sample)
		sum += lastPow
		if sum > pos {
			return i
		}
	}
	return 0
}
