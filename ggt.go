package main

import (
	"fmt"
)

func GGT(a int, b int) int {
	if b == 0 {
	return a
	}else{return GGT(b, a % b)}
}

func main() {
	fmt.Println(GGT(1234, 5))
}
