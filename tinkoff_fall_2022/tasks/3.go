package tasks

import (
	"fmt"
)

func T3() {
	var count int
	fmt.Scanf("%d", &count)

	subs := make([]int, (count/2)+count%2)
	cans := make([]int, count/2)
	flag := true
	j := 0
	for i := 0; i < count; i++ {
		if flag {
			fmt.Scanf("%d", &subs[j])
		} else {
			fmt.Scanf("%d", &cans[j])
			j++
		}

		flag = !flag
	}

	var min int
	for i := range subs {
		if subs[i] < subs[min] {
			min = i
		}
	}

	var max int
	for i := range cans {
		if cans[i] > cans[max] {
			max = i
		}
	}

	if cans[max] > subs[min] {
		temp := cans[max]
		cans[max] = subs[min]
		subs[min] = temp
	}

	sum := 0
	for i := range subs {
		sum += subs[i]
	}

	for i := range cans {
		sum -= cans[i]
	}

	fmt.Println(sum)
}
