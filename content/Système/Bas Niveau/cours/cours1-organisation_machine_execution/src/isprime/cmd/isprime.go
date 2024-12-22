package main

import (
	"math"
)

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19}

func IsPrime(num int) bool {
	if len(primes) == 8 {
		initPrimes()
	}
	if num < primes[len(primes)-1] {
		return numIsInPrimes(num, 0, len(primes)-1)
	}
	return numIsPrime(num)
}

func initPrimes() {
	for i := 21; i < 1000; i += 2 {
		sqrt_i := int(math.Ceil(math.Sqrt(float64(i))))
		prime := true
		for j := 3; j <= sqrt_i; j += 2 {
			if i%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			// log.Printf("adding %d to primes\n", i)
			primes = append(primes, i)
		}
	}
}

func numIsInPrimes(num, lo, hi int) bool {
	if lo == hi {
		return num == primes[lo]
	}
	mid := lo + (hi-lo)/2
	if primes[mid] == num {
		return true
	}
	if primes[mid] < num {
		return numIsInPrimes(num, mid+1, hi)
	} else {
		return numIsInPrimes(num, lo, mid)
	}
}

func numIsPrime(num int) bool {
	sqrt_num := int(math.Ceil(math.Sqrt(float64(num))))
	if sqrt_num > primes[len(primes)-1] {
		for i := 0; i < len(primes); i++ {
			if num%primes[i] == 0 {
				return false
			}
		}
		for j := primes[len(primes)-1] + 1; j <= sqrt_num; j++ {
			if num%j == 0 {
				return false
			}
		}
	} else {
		for i := 0; primes[i] <= sqrt_num; i++ {
			if num%primes[i] == 0 {
				return false
			}
		}
	}
	return true
}
