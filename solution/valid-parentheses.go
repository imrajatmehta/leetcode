package solution

import "github.com/imrajatmehta/leetcode/library"

func IsValid(s string) bool {
	stack := library.New()
	for i := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack.Push(s[i])
		} else if s[i] == '}' && stack.Peek().(byte) == '{' {
			stack.Pop()
		} else if s[i] == ']' && stack.Peek().(byte) == '[' {
			stack.Pop()
		} else if s[i] == ')' && stack.Peek().(byte) == '(' {
			stack.Pop()
		} else {
			return false
		}
	}
	return stack.Len() <= 0
}
