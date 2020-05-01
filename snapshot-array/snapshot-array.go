package main

func main() {

}

type SnapshotArray struct {
	metadata [][][]int
	seq      int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		metadata: make([][][]int, length),
		seq:      0,
	}
}
func (sa *SnapshotArray) Set(index int, val int) {
	if len(sa.metadata[index]) == 0 || sa.metadata[index][len(sa.metadata[index])-1][0] != sa.seq {
		sa.metadata[index] = append(sa.metadata[index], []int{sa.seq, val})
		return
	}
	sa.metadata[index][len(sa.metadata[index])-1][1] = val
}

func (sa *SnapshotArray) Snap() int {
	sa.seq++
	return sa.seq - 1
}

func (sa *SnapshotArray) Get(index int, snapID int) int {
	return findLE(sa.metadata[index], snapID)
}

func findLE(data [][]int, snapID int) int {
	if len(data) == 0 {
		return 0
	}
	if data[0][0] > snapID {
		return 0
	}
	if data[len(data)-1][0] <= snapID {
		return data[len(data)-1][1]
	}
	l, r := 0, len(data)-1
	for l <= r {
		m := (l + r) >> 1
		if data[m][0] == snapID {
			return data[m][1]
		} else if data[m][0] < snapID {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return data[r][1]
}
