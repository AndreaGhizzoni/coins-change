package main

/*
	What is the minimum number of coins to give back as change ?
	Input: the change N and a set of coins cut T.
	Output: a list of coins cut that summed equals to N. That list must
			contains the minimum number of elements.
*/

// This solution use a greedy approach.

import (
	"fmt"
	"sort"
)

// This function takes a slice of coins cut T and the change N and return a
// slice that contains the number of coins cut to give as change.
// The greedy decision in this case is: always choose the highest cut of coin
// that mod N == 0.
// In order to make the greedy decision, this method needs to reverse sort the
// given slice T such that: T[0] > T[1] > T[2] > ... > T[len(T)-1]. To do that
// it creates a copy or T and reverse sort it (because sort package sort the
// given slice and not produce a sorted copy).
// Complexity: O( n + n*log(n) + n ) = O( n*log(n) )
func change(T []int, N int) []int {
	rT := make([]int, len(T))
	copy(rT, T)
	sort.Sort(sort.Reverse(sort.IntSlice(rT)))

	A := make([]int, len(rT))
	for i, rTi := range rT {
		A[i] = N / rTi
		N = N - A[i]*rTi
	}
	return A
}

// This function takes a slice of coins cut T and the change N and print a list
// of coins cut to give as a change.
func printChange(T []int, N int) {
	S := change(T, N)

	lenT := len(T) - 1
	for i := len(S) - 1; i >= 0; i-- {
		for j := 0; j < S[i]; j++ {
			// lenT-i because S is calculated based on sort.Reverse(T)
			fmt.Println(T[lenT-i])
		}
	}
}

func main() {
	// size of coins.
	T := []int{1, 2, 5, 10, 20, 50, 100, 200}
	// change amount
	N := 27

	fmt.Println("With the following cut of coins:")
	fmt.Println(T)
	fmt.Println("And change as:", N)
	fmt.Println("The minimum number of coins cut to give as a change are:")
	printChange(T, N)
}
