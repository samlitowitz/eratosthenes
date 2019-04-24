# Sieve of Eratosthenes
[![Go Report Card](https://goreportcard.com/badge/github.com/samlitowitz/eratosthenes)](https://goreportcard.com/report/github.com/samlitowitz/eratosthenes)
[![GoDoc](https://godoc.org/github.com/samlitowitz/eratosthenes/pkg/sieve?status.svg)](https://godoc.org/github.com/samlitowitz/eratosthenes/pkg/sieve)
[![Build Status](https://travis-ci.org/samlitowitz/eratosthenes.svg?branch=master)](https://travis-ci.org/samlitowitz/eratosthenes)

Read about the [Sieve of Eratosthenes](https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes).

## Building the command line tool
```
git clone https://github.com/samlitowitz/eratosthenes.git
cd eratosthenes
go build -o sieve cmd/sieve
```

## Using the command line tool
```
Usage of sieve-of-eratosthenes:
  -max int
    	max integer (default 64)
```

## Using the package
```go
package main

import (
	"fmt"
	"github.com/samlitowitz/eratosthenes/pkg/sieve"
)

func main() {
	naive, _ := sieve.NewNaive(64)
	primes := naive.Primes()
	fmt.Printf("%+v\n", primes)
	
	segmented, _ := sieve.NewSegmented(64, 64)
	primes = segmented.Primes()
	fmt.Printf("%+v\n", primes)	
}

```

## License
Apache 2.0
