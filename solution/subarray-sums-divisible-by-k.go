package solution

func SubarraysDivByK(nums []int, k int) int {
	tot := 0
	freq := make(map[int]int)
	pSum := make([]int, len(nums))

	for i := range nums {
		if i > 0 {
			pSum[i] = pSum[i-1]
		}
		pSum[i] += nums[i] % k
	}
	for i := range nums {
		if val, ok := freq[pSum[i]-k]; ok {
			tot += val
		}
		if val, ok := freq[-pSum[i]]; ok {
			tot += val
		}
		if val, ok := freq[pSum[i]+k]; ok {
			tot += val
		}
		if pSum[i]%k == 0 {
			tot++
		}
		freq[pSum[i]]++
	}
	return tot
}
