package solution
func LongestSubsequence(arr []int, diff int) int {
    hash:=make(map[int]int)
    globalMax:=1
    for i:=range arr{
        val1:=0
        val2:=0
        val1,_=hash[arr[i]-diff]
        val2,_=hash[arr[i]]
        max:=Max(1+val1,val2)
        hash[arr[i]]=max
        globalMax=Max(globalMax,max)
    }
    return globalMax
}
func Max(a,b int)int{
    if a>b{
        return a
    }
    return b
}
