package pool

import (
	"crypto/md5"
	"testing"
)

var sample []rune = []rune("abcdefghijklmnopqrstuvwxyz")

func Benchmark_CrackDefault(b *testing.B) {
	var hash [16]byte = md5.Sum([]byte("psswrd"))
	for i := 0; i < b.N; i++ {
		p := New(sample, func(pwd []byte) bool {
			return md5.Sum(pwd) == hash
		})
		p.SetLogging(false)
		p.Crack()
	}
}

func Test_Crack_CanCrackPasswords(t *testing.T) {
	passwords := []string{
		"asdf",
		"kdal",
		"qzyp",
		"adgae",
		"oljp",
	}

	for _, password := range passwords {
		hash := md5.Sum([]byte(password))
		p := New(sample, func(pwd []byte) bool {
			return md5.Sum(pwd) == hash
		})
		p.SetLogging(false)
		p.Crack()
		p.SetLogging(false)
		crackedPassword := p.Crack()

		if crackedPassword != password {
			t.Errorf("Password was \"%s\" but pool.Crack returned: \"%s\" \n", password, crackedPassword)
		}

	}
}
