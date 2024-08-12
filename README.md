# Gocrack

Gocrack is a simple hash cracker for md5 hashes.
It checks all strings possible with letters from the given sample.
If no sample is given the default is a to z in all lowercase.

Gocrack is intended to be be a training project for myself to become
more familiar with go's concurrency & parallelism features.  

## Usage

Here is an example:

``` sh
gocrack -sample=abcdeGhiyklNmr 94be2ede1a514fc0c5647abed54b8c7f
```

To run gocrack you have to give it a md5 hash in hex as an argument.
If you want more control over the performance of gocrack you can
use following flags:

```
  -blocksize int
    	 (default 5000)
  -sample string
    	 (default "abcdefghijklmnopqrstuvwxyz")
  -workers int
```

## Installation


``` sh
git clone https://github.com/chlyNiklas/gocrack.git
cd gocrack
go install .
```

