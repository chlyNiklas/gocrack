// checkfunctionfactory provides everything needed to create a CheckFunction.
package checkfunctionfactory

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
)

// CheckFunction takes a byte-slice and check's if it is the input
// that when hashed is the hash.
type CheckFunction func([]byte) bool

// HashAlgorithm is a enum for HashAlgorithms
type HashAlgorithm string

const (
	Sha256 HashAlgorithm = "sha256"
	Sha512               = "sha512"
	Sha1                 = "sha1"
	Md5                  = "md5"
)

// CheckFunctionFactory builds CheckFunctions.
type CheckFunctionFactory struct {
	algorithm HashAlgorithm
	hashSlice []byte
}

func New(algorithm HashAlgorithm, hashString string) (factory *CheckFunctionFactory, err error) {

	hashSlice, err := hex.DecodeString(hashString)
	if err != nil {
		return
	}

	return &CheckFunctionFactory{
		algorithm: algorithm,
		hashSlice: hashSlice,
	}, nil
}

// CreateCheckFunction returns a check function for the set algorithm.
// If the hashString is invalid or the algorithm is invalid / not implemented
// CreateCheckFunction returns an error.
func (c *CheckFunctionFactory) CreateCheckFunction() (function CheckFunction, err error) {
	switch c.algorithm {
	case Sha256:
		hash := [sha256.Size]byte(c.hashSlice)

		function = func(b []byte) bool {
			return sha256.Sum256(b) == hash
		}
	case Sha512:
		hash := [sha512.Size]byte(c.hashSlice)

		function = func(b []byte) bool {
			return sha512.Sum512(b) == hash
		}
	case Sha1:
		hash := [sha1.Size]byte(c.hashSlice)

		function = func(b []byte) bool {
			return sha1.Sum(b) == hash
		}
	case Md5:
		hash := [md5.Size]byte(c.hashSlice)

		function = func(b []byte) bool {
			return md5.Sum(b) == hash
		}

	default:
		return nil, errors.New("Hash function not implemented.")
	}

	return
}
