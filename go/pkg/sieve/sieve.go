package sieve

import (
	"fmt"
	"math"
)

type Sieve interface {
	NthPrime(n int64) int64
}

// SieveData saves primes so multiple NthPrime calls can re-use the primes already determined
type SieveData struct {
	primes []int64
}

func NewSieve() Sieve {
	return &SieveData{
		[]int64{2},
	}
}

func (s *SieveData) NthPrime(n int64) int64 {
	if n < 0 {
		// I prefer logging to panicking for easier debugging...
		// (if the interface supported it, I'd return an error)
		fmt.Println("failure - n cannot be negative")
		return 0
	} else if n < int64(len(s.primes)) { // use cache
		return s.primes[n]
	}

	if n < 6 {
		s.sieveToCap(10_000)
	} else {
		// per [Wikipedia](https://en.wikipedia.org/wiki/Prime_number_theorem), this function
		// gives an upper bound for the nth prime for all n >= 6
		s.sieveToCap(int64(math.Ceil(float64(n) * (math.Log(float64(n)) + math.Log(math.Log(float64(n)))))))
	}

	if int64(len(s.primes)) > n { // _should_ always be true
		return s.primes[n]
	} else {
		fmt.Println("failure - did not generate enough primes")
		return 0
	}
}

// sieveToCap runs the sieve of Eratosthenes to generate all primes up to a value (cap)
func (s *SieveData) sieveToCap(cap int64) {
	var isPrime = make([]bool, cap+1)
	for i := range isPrime {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false

	// Run the known primes through the sieve
	for _, prime := range s.primes {
		for i := prime * prime; i <= cap; i += prime {
			isPrime[i] = false
		}
	}

	//  then start catching new primes
	for i := s.primes[len(s.primes)-1] + 1; i <= cap; i++ {
		if isPrime[i] {
			nextPrime := i
			s.primes = append(s.primes, nextPrime)
			for i := nextPrime * nextPrime; i <= cap; i += nextPrime {
				isPrime[i] = false
			}
		}
	}
}
