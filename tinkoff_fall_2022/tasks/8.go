package tasks

import "fmt"

func startsWith(s, start string) bool {
	return s[:len(start)] == start
}

func endsWith(s, end string) bool {
	return s[len(s)-len(end):] == end // WARNING
}

func T8() {
	var p, q int
	fmt.Scanf("%d %d", &p, &q)

	domains := make([]string, p)
	for i := range domains {
		fmt.Scanf("%s", &domains[i])
	}

	type order struct {
		Start, End string
	}
	orders := make([]order, q)
	for i := range orders {
		fmt.Scanf("%s %s", &orders[i].Start, &orders[i].End)
	}

	for _, o := range orders {
		count := 0
		for _, d := range domains {
			if startsWith(d, o.Start) && endsWith(d, o.End) {
				count++
			}
		}
		fmt.Println(count)
	}
}
