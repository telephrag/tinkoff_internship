package tasks

import (
	"fmt"
	"sort"
)

func T1() {
	x := make([]int, 4)
	y := make([]int, 4)

	fmt.Scanf("%d", &x[0])
	fmt.Scanf("%d", &y[0])
	fmt.Scanf("%d", &x[1])
	fmt.Scanf("%d", &y[1])

	fmt.Scanf("%d", &x[2])
	fmt.Scanf("%d", &y[2])
	fmt.Scanf("%d", &x[3])
	fmt.Scanf("%d", &y[3])

	sort.Ints(x)
	sort.Ints(y)

	width := x[0] - x[3]
	if width < 0 {
		width *= -1
	}
	height := y[0] - y[3]
	if height < 0 {
		height *= -1
	}

	if width > height {
		fmt.Println(width * width)
	} else {
		fmt.Println(height * height)
	}
}
