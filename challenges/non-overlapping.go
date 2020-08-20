package challenges

import "sort"

// Interval - struct to mainitain starting and ending points
type Interval struct {
	start int
	end   int
}
type intervals []Interval

func (interval intervals) Less(x, y int) bool {
	return interval[x].end < interval[y].end
}
func (interval intervals) Len() int {
	return len(interval)
}

func (interval intervals) Swap(i, j int) {
	interval[i], interval[j] = interval[j], interval[i]
}

func eraseOverlapIntervals(intervals [][]int) int {
	intervalsStruct := getIntervals(intervals)

	if len(intervalsStruct) == 0 {
		return 0
	}

	sort.Sort(intervalsStruct)
	count := 1
	start := intervalsStruct[0].end
	for i := 1; i < len(intervalsStruct); i++ {
		if intervalsStruct[i].start < start {
			continue
		}
		count++
		start = intervalsStruct[i].end
	}

	return len(intervalsStruct) - count

}

func getIntervals(intervals [][]int) intervals {
	var intervalsStructs []Interval
	for _, ele := range intervals {
		intervalsStructs = append(intervalsStructs, Interval{ele[0], ele[1]})
	}
	return intervalsStructs
}
