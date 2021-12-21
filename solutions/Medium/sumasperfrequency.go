/*
Problem Statement
There are  numbers N, and you are given  queries. In each query, you are given two integers l and r.

You are required to print the sum of all the numbers whose frequency of occurrence is between  and  (including l and r).
Print a single integer for each query in a new line.
*/
package main

import (
	"fmt"
)

func main() {
	//array size
	var size int
	//query size
	var qsize int

	//map to store frequency of the number occured
	f := make(map[int]int)
	fmt.Scanln(&size)
	//array of size(size) to store elements
	var arr = make([]int, size)
	for i := 0; i < size; i++ {

		fmt.Scanf("%d", &arr[i])
		f[arr[i]] = f[arr[i]] + 1
	}

	fmt.Scanln(&qsize)
	//array to store query
	var earray [2]int
	sumarr := make([]int, qsize)
	for i := 0; i < qsize; i++ {
		sum := 0
		for j := 0; j < 2; j++ {

			fmt.Scanf("%d", &earray[j])
			for s := range f {
				if earray[j] <= f[s] || f[s] >= earray[j] {
					sum = sum + s*f[s]
				}
			}
			sumarr[i] = sum
		}

	}
	//print to stdout
	for _, v := range sumarr {
		fmt.Println(v)
	}

}
