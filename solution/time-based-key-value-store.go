package solution

type TimeMap struct {
	Map map[string][]PairTV
}
type PairTV struct {
	Timestamp int
	Value     string
}

func Constructor() TimeMap {
	var obj TimeMap
	obj.Map = make(map[string][]PairTV)
	return obj

}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.Map[key]; !ok {
		this.Map[key] = []PairTV{}
	}
	this.Map[key] = append(this.Map[key], PairTV{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {

	ind := binarySearchOnPairTV(this.Map[key], timestamp)
	if ind == -1 {
		return ""

	}
	return this.Map[key][ind].Value

}
func binarySearchOnPairTV(arr []PairTV, timestamp int) int {
	left := 0
	right := len(arr) - 1
	ans := -1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid].Timestamp <= timestamp {
			ans = mid
			left = mid + 1

		} else {
			right = mid - 1
		}
	}
	return ans
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
