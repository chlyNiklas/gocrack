# Gocrack

Gocrack is a simple hash cracker for md5 hashes.
It checks all strings possible with letters from the given charset.
If no charset is given the default is a to z in all lowercase.

Gocrack is intended to be be a training project for myself to become
more familiar with go's concurrency & parallelism features.  

## Usage

Here is an example:

``` sh
gocrack -hashType md5 -charset=abcdeGhiyklNmr 94be2ede1a514fc0c5647abed54b8c7f
```

To run gocrack you have to give it a hash in hex as an argument.
If you want more control over the performance of gocrack you can
use following flags:

```
-blocksize int
    Set the number of hashes one worker checks before waiting for a new job
-charset string
    All characters your password could possibly contain (default "abcdefghijklmnopqrstuvwxyz")
-hashType string
    The hash algorithm used to make your hash. You can choose between md5 & sha256 (default "md5")
-workers int
    Set the number of go routines that crack
```

## Installation


``` sh
git clone https://github.com/chlyNiklas/gocrack.git
cd gocrack
go install .
```

