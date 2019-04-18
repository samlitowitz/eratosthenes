package main

import (
	"flag"
	"fmt"
	"github.com/samlitowitz/eratosthenes/pkg/sieve"
	"log"
)

func main() {
	var n int64

	flag.Int64Var(&n, "n", 64, "upper bound of search range")
	flag.Parse()

	fmt.Printf("Finding all primes in [2, %d]\n", n)

	naive, err := sieve.NewNaive(n)
	if err != nil {
		log.Fatal(err)
	}

	primes := naive.Primes()

	fmt.Printf("%+v\n", primes)
}
