package main

/*
intervals = [[1,2], [3,5],[6,7],[8,10],[12,16]]
newInterval = [4,8]

toAdd = [3,8]

intervals = [[1,2], [3,8],[6,7],[8,10],[12,16]]
intervals = [[1,2], [3,10],[12,16]]
newInterval = [4,8]

*/
func insert(intervals [][]int, newInterval []int) [][]int {

	res := [][]int{}
	toAdd := newInterval

	for i := 0; i < len(intervals); i++ {

		if intervals[i][0] > toAdd[1] { //no overlap addTO is before
			res = append(res, toAdd)
			toAdd = intervals[i]
		} else if intervals[i][1] >= toAdd[0] { //has overlap
			toAdd[0] = Min(intervals[i][0], toAdd[0])
			toAdd[1] = Max(intervals[i][1], toAdd[1])
		} else {
			res = append(res, intervals[i])
		}
	}

	res = append(res, toAdd)

	return res
}

func Max(a, b int) int {

	if a < b {
		return b
	}

	return a
}

func Min(a, b int) int {

	if a > b {
		return b
	}

	return a

}

/*

`       `       0.    1    2      3.     4
intervals = [[1,2], [3,5],[6,7],[8,10],[12,16]]
newInterval = [4,8]

toAdd = [3,8]

intervals = [[1,2], [3,8],[6,7],[8,10],[12,16]]
intervals = [[1,2], [3,10],[12,16]]
newInterval = [4,8]

1:  figure out where to put my new interval - iterate  through intervals
2: merge inrvals


1:
newstart =newInterval[0]
newEnd = newInterval[1]
for i -- range intervals
    start := intervals[i][0]
    end := intervals[i][1]

    if newstart <= end {
        min:= Min(start, newStart)
        max:= Max(end, newEnd)
        intervals[i][0] = min
        intervals[i][1] = max
    }ese{
        continue
    }
}

res := [][]int{}

prevStart := intervals[0][0]
prevEnd := intervals[0][1]
for 1 -- range intervals
    start := intervals[i][0]
    end := intervals[i][1]

    if start <=  prevEnd {

        if end > prevEnd {
            res = append(res,[]int{start,end})
        }

    }else{

        res = append(res,intervals[i])
    }


*/
