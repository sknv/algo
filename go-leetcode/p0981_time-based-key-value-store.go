package leetcode

// https://leetcode.com/problems/time-based-key-value-store/description/

type TimeMap struct {
	hash map[string]timestamps
}

// TimeMapConstructor = Constructor
func TimeMapConstructor() TimeMap {
	return TimeMap{
		hash: make(map[string]timestamps),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.hash[key] = append(this.hash[key], timestampEntry{
		timestamp: timestamp,
		value:     value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	timestamps := this.hash[key]

	return timestamps.search(timestamp)
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

type timestampEntry struct {
	timestamp int
	value     string
}

type timestamps []timestampEntry

func (ts timestamps) search(timestamp int) string {
	if len(ts) == 0 {
		return ""
	}

	// Можно было использовать стандартную функцию sort.Search, но мы напишем свой бинарный поиск.
	left, right := 0, len(ts)-1
	for left <= right {
		mid := left + (right-left)/2
		currentTimestamp := ts[mid].timestamp

		if currentTimestamp == timestamp {
			return ts[mid].value
		}

		if timestamp < currentTimestamp {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// Если нет промежуточных значений, то по условию задачи возвращаем пустую строку.
	if right < 0 {
		return ""
	}

	return ts[right].value
}
