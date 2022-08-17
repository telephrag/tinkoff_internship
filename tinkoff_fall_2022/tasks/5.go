package tasks

import (
	"fmt"
	"sort"
)

func T5() {
	var n, q int
	fmt.Scanf("%d %d", &n, &q)

	names := make([]string, n)
	queriesStr := make([]string, q)
	queriesPos := make([]int, q)

	for i := range names {
		fmt.Scanf("%s", &names[i])
	}

	for i := range queriesStr {
		fmt.Scanf("%d", &queriesPos[i])
		fmt.Scanf("%s", &queriesStr[i])
	}

	for i, prefix := range queriesStr {
		hasPrefix := make([]string, 0)
		l := len(prefix)
		for _, name := range names {
			if len(name) >= l { // bounds check
				if name[:l] == prefix {
					hasPrefix = append(hasPrefix, name)
				}
			}
		}

		if len(hasPrefix) < queriesPos[i] {
			fmt.Println(-1)
			continue
		}

		sort.Strings(hasPrefix)
		for j := range names {
			if names[j] == hasPrefix[queriesPos[i]-1] {
				fmt.Println(j + 1)
				break
			}
		}
	}
}
