func minFlips(a int, b int, c int) int {
    ans:=0
    for i:=0;i<=30;i++{
        aa:=a&1
        bb:=b&1
        cc:=c&1
        // fmt.Println(aa,bb,cc)
        if cc==0{
            ans+=aa+bb
        }else{
            if aa==0 && bb==0{
                ans++
            }
        }

        a/=2
        b/=2
        c/=2
    }    
    return ans
}
