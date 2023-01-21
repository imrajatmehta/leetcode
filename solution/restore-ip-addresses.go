package solution

import (
	"strconv"
	"strings"
)

func RestoreIpAddresses(s string) []string {
	ans := []string{}
	putDotInStrings(0, &ans, s, "")
	return ans
}
func putDotInStrings(i int, ans *[]string, s, ip string) {
	if i >= len(s) {
		if isValidIp(ip) {
			(*ans) = append((*ans), ip)
		}
		return
	}
	putDotInStrings(i+1, ans, s, ip+string([]byte{s[i]}))
	if i > 0 {
		putDotInStrings(i+1, ans, s, ip+"."+string([]byte{s[i]}))
	}
}

func isValidIpNum(num string) bool {

	if len(num) > 1 && num[0] == '0' {
		return false
	}
	a, err := strconv.Atoi(num)
	if err != nil {
		return false
	}
	return a >= 0 && a <= 255
}
func isValidIp(ip string) bool {
	ips := strings.Split(ip, ".")

	if len(ips) != 4 {
		return false
	}
	// fmt.Println(ips)
	for i := range ips {
		if !isValidIpNum(ips[i]) {
			return false
		}
	}
	return true
}
