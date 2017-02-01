package main

/*
	What is the minimum number of coins to give back as change ?
	Input: the change N and a set of coins cut T.
	Output: a list of coins cut that summed equals to N. That list must
			contains the minimum number of elements.
*/

// This solution use a dynamic programming approach.

import (
	"fmt"
	"math"
)

// This function takes a slice of coins cuts T and the amount of change N.
// Returns a slice of choices of T.
// Complexity: O( N*len(T) ) ~ O( N^2 )
func change(T []int, N int) []int {
	// coins choices. N+1 because we will use S[1_N]
	S := make([]int, N+1)
	// auxiliary slice. N+1 because we will use C[1_N]
	C := make([]int, N+1)
	C[0] = 0

	// solve the problem N bottom-up. m is the current change.
	for m := 1; m <= N; m++ {
		C[m] = math.MaxInt32 // set current change to max

		// for every coins cut, check if I can give back that cut
		for j := 0; j < len(T); j++ {
			if m >= T[j] && C[m-T[j]]+1 < C[m] {
				C[m] = C[m-T[j]] + 1
				// save that j+1 can be the change for the problem m
				S[m] = j + 1
			}
		}
	}

	return S
}

// This function takes a slice of coins cuts T, the change to give back N
// and print a list of coins cut, taken from T, to give the change of N.
func printChange(T []int, N int) {
	S := change(T, N)

	for N > 0 {
		fmt.Println(T[S[N]-1])
		N = N - T[S[N]-1]
	}
}

func main() {
	// size of coins
	T := []int{1, 2, 5, 10, 20, 50, 100, 200}
	// change amount
	N := 27

	fmt.Println("With the following cut of coins:")
	fmt.Println(T)
	fmt.Println("And change as: ", N)
	fmt.Println("The minimum number of coins cut to give as a change are:")
	printChange(T, N)
}
