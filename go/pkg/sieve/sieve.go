package sieve

import "fmt"

type Sieve interface {
	NthPrime(n int64) int64
}

// SieveData saves primes so multiple NthPrime calls can re-use the primes already determined
type SieveData struct {
	primes []int64
}

func NewSieve() Sieve {
	return SieveData{
		[]int64{2},
	}
}

func (s SieveData) NthPrime(n int64) int64 {
	// TODO: more intelligent cap
	s.sieveToCap(10_000)

	if int64(len(s.primes)) > n { // _should_ always be true
		return s.primes[n]
	} else {
		fmt.Println("failure - did not generate enough primes") // prefer logging to panicking for easier debugging...
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
	for i := int64(3); i <= cap; i++ {
		if isPrime[i] {
			nextPrime := i
			s.primes = append(s.primes, nextPrime)
			for i := nextPrime * nextPrime; i <= cap; i += nextPrime {
				isPrime[i] = false
			}
		}
	}
}
