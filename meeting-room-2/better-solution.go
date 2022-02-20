func minMeetingRooms(intervals [][]int) int {

	if len(intervals) == 0 {
		return 0
	}

	start, end := make([]int, len(intervals)), make([]int, len(intervals))

	for i := 0; i < len(intervals); i++ {
		start[i] = intervals[i][0]
		end[i] = intervals[i][1]
	}

	sort.Ints(start)
	sort.Ints(end)

	meetingRoom, preIndex := 0, 0

	for i := 0; i < len(intervals); i++ {

		if start[i] < end[preIndex] {
			meetingRoom++
		} else {
			preIndex++
		}
	}

	return meetingRoom

}

