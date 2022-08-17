package tasks

import (
	"fmt"
	"sort"
)

func T2() {
	var count int
	fmt.Scanf("%d", &count)

	teams := make([][]string, count)
	for i := 0; i < count; i++ {
		teams[i] = make([]string, 3)
		fmt.Scanf("%s %s %s", &teams[i][0], &teams[i][1], &teams[i][2])
		sort.Strings(teams[i])
	}

	maxWins := 1
	for i := range teams {
		this := 1
		for j := i + 1; j < len(teams); j++ {
			if (teams[j][0] + teams[j][1] + teams[j][2]) == (teams[i][0] + teams[i][1] + teams[i][2]) {
				this++
			}
		}

		if this > maxWins {
			maxWins = this
		}
	}

	fmt.Println(maxWins)
}
