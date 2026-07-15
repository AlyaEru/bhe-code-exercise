package sieve

import (
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNthPrime(t *testing.T) {
	t.Run("timing tests", func(t *testing.T) {
		// large prime doesn't take too long
		// (timing could vary depending where this test is running,
		// but we want _some_ kind of benchmarking)
		sieve := NewSieve()

		start := time.Now()
		sieve.NthPrime(10_000_000)
		elapsed := time.Since(start)
		assert.Less(t, elapsed, time.Second*3)

		// Sieve caches primes: running on smaller n is very fast
		start = time.Now()
		sieve.NthPrime(9_999_999)
		elapsed = time.Since(start)
		assert.Less(t, elapsed, time.Millisecond*100)

	})

	t.Run("value tests", func(t *testing.T) {
		cases := []struct {
			input    int64
			expected int64
		}{
			{input: -6, expected: 0},
			{input: 0, expected: 2},
			{input: 1, expected: 3},
			{input: 19, expected: 71},
			{input: 99, expected: 541},
			{input: 500, expected: 3_581},
			{input: 986, expected: 7_793},
			{input: 2_000, expected: 17_393},
			{input: 1_000_000, expected: 15_485_867},
			{input: 10_000_000, expected: 179_424_691},
			// { input: 100_000_000, expected: 2_038_074_751}, // this works but takes a bit
		}

		sieve := NewSieve()
		for _, cc := range cases {
			assert.Equal(t, cc.expected, sieve.NthPrime(cc.input))
		}
	})
}

func FuzzNthPrime(f *testing.F) {
	sieve := NewSieve()

	f.Fuzz(func(t *testing.T, n int64) {
		if !big.NewInt(sieve.NthPrime(n)).ProbablyPrime(0) {
			t.Errorf("the sieve produced a non-prime number at index %d", n)
		}
	})
}
