package solution

type UndergroundSystem struct {
	TimeMap map[string][]int
	Train   map[int]UndergroundSystemTrain
}

type UndergroundSystemTrain struct {
	Station string
	Time    int
}

func UndergroundSystemConstructor() UndergroundSystem {
	temo := UndergroundSystem{}
	temo.TimeMap = make(map[string][]int)
	temo.Train = make(map[int]UndergroundSystemTrain)
	return temo

}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.Train[id] = UndergroundSystemTrain{Station: stationName, Time: t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, endTime int) {
	key := this.Train[id].Station + "=" + stationName
	startTime := this.Train[id].Time
	if _, ok := this.TimeMap[key]; ok {
		this.TimeMap[key][0] += endTime - startTime
		this.TimeMap[key][1] += 1
	} else {
		this.TimeMap[key] = []int{endTime - startTime, 1}
	}
}
func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	aa := float64(this.TimeMap[startStation+"="+endStation][1])
	bb := float64(this.TimeMap[startStation+"="+endStation][0])
	return bb / aa

}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
