package main

import "fmt"

func numberofsteps() {

	//array size
	var size int

	fmt.Scanln(&size)

	//array 'a' of size 'size'

	a := make([]int, size)

	for i := 0; i < size; i++ {

		fmt.Scanf("%d", &a[i])
	}

	min := a[0]

	//array 'b' of size 'size'

	b := make([]int, size)

	for i := 0; i < size; i++ {

		fmt.Scanf("%d", &b[i])
	}
	count := 0
	i := 0
	for i < size {

		for a[i] > min {
			a[i] = a[i] - b[i]
			count = count + 1
		}
		if a[i] <= min {
			min = a[i]
			i = 0
		} else if a[i] < 0 {
			count = 1
			break
		}
		i = i + 1
	}

	fmt.Println(count)
}

// func minimum(a []int) int {

// 	min := a[0]

// 	for i := 0; i < len(a); i++ {

// 		if a[i] <= min {
// 			min = a[i]
// 		}
// 	}

// 	return min

// }
