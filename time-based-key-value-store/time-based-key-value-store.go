package main

type TimeMap struct {
	TimeValues map[string][]ValueStamp
}

type ValueStamp struct {
	Value string
	Time  int
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{
		TimeValues: make(map[string][]ValueStamp),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	var timeStamp ValueStamp
	timeStamp.Value = value
	timeStamp.Time = timestamp

	this.TimeValues[key] = append(this.TimeValues[key], timeStamp)
}

func (this *TimeMap) Get(key string, timestamp int) string {

	if len(this.TimeValues[key]) == 0 || this.TimeValues[key][0].Time > timestamp {
		return ""
	}

	valArr := this.TimeValues[key]

	left := 0
	right := len(valArr) - 1

	for left <= right {
		mid := (left + right) / 2
		if timestamp < valArr[mid].Time {
			right = mid - 1

		} else {
			left = mid + 1
		}

	}
	return valArr[right].Value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
