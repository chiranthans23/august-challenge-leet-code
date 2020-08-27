package challenges

import (
	"fmt"
	"sort"
)

// Period -
type Period struct {
	start int
	end   int
}
type periodArray []Period

func (s periodArray) Len() int {
	return len(s)
}
func (s periodArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s periodArray) Less(i, j int) bool {
	return s[i].start <= s[j].start
}

func findRightInterval(intervals [][]int) []int {
	p := getPeriodAsStruct(intervals)
	sort.Sort(periodArray(p))
	hasRight := make(map[Period]int)
	indices := make(map[Period]int)
	for i, ele := range intervals {
		indices[Period{ele[0], ele[1]}] = i
	}
	fmt.Printf("Indices array: %v\n", indices)
	for i := 0; i < len(p); i++ {
		cur := sort.Search(len(p), func(j int) bool { return p[j].start >= p[i].end })
		if cur == len(p) {
			hasRight[p[i]] = -1
		} else {
			hasRight[p[i]] = indices[p[cur]]
		}
	}
	var rightInterval []int
	for _, ele := range intervals {
		rightInterval = append(rightInterval, hasRight[Period{ele[0], ele[1]}])
	}
	return rightInterval
}

func getPeriodAsStruct(intervals [][]int) []Period {
	var p []Period
	for _, ele := range intervals {
		p = append(p, Period{ele[0], ele[1]})
	}
	return p
}
func getIndexofNextInterval(ele Period, point int, p []Period, indices map[Period]int) int {
	for i := 0; i < len(p)-1; i++ {
		if p[i].end < p[i+1].start {
			return indices[p[i+1]]
		}
	}
	return -1
}
