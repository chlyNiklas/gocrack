package solver

import "github.com/chlyniklas/gocrack/checkfunctionfactory"

type Solver struct {
	// the sample is a byte-slice so that the program has to do less conversions between datatype
	sample        []byte
	checkFunction checkfunctionfactory.CheckFunction
}

func New(sample []byte, checkFunction checkfunctionfactory.CheckFunction) *Solver {
	return &Solver{
		sample,
		checkFunction,
	}
}

func (s *Solver) CheckCombinationAtPosition(pos int) (str []byte, ok bool) {
	str = s.CreateUniqueCombination(pos)
	if s.checkFunction(s.CreateUniqueCombination(pos)) {
		return str, true
	}
	return str, false
}

func (s *Solver) CreateUniqueCombination(pos int) []byte {
	nDigits := s.digits(pos)

	pwd := make([]byte, 0, nDigits)
	for range nDigits {
		pwd = append(pwd, s.sample[pos%len(s.sample)])
		pos = pos / len(s.sample)
	}
	return pwd
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
