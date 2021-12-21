package main

import (
	"fmt"
	"strings"
)

func splithouse() {

	var size int
	fmt.Scanln(&size)
	var grid string

	fmt.Scanf("%s", &grid)

	s := strings.ReplaceAll(grid, ".", "B")
	if !strings.Contains(s, "HH") {
		fmt.Printf("YES\n%s", s)

	}
	fmt.Println("NO")

}
