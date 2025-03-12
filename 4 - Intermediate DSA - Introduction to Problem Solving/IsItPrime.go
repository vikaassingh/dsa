package dsa

import (
	"fmt"
)

func Prime() {
	var A int
	fmt.Scanf("%d", &A)
	flag := true
	for i := 2; i < A/2; i++ {
		if A%i == 0 {
			flag = false
			break
		}
	}
	if flag {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
