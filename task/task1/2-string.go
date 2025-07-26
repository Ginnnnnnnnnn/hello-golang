package main

import "fmt"

func main() {
	// 有效的括号
	isValidDatas := []string{"()", "()[]{}", "(]", "([])", "([)]"}
	for _, value := range isValidDatas {
		result := isValid(value)
		fmt.Println(result)
	}
	// 最长公共前缀
	longestCommonPrefixDatas := [][]string{{"flower", "flow", "flight"}, {"dog", "racecar", "car"}}
	for _, value := range longestCommonPrefixDatas {
		result := longestCommonPrefix(value)
		fmt.Println(result)
	}
}

func isValid(s string) bool {
	var stack []string
	for _, value := range s {
		str := string(value)
		if len(stack) == 0 {
			stack = append(stack, str)
		} else {
			if str == ")" {
				if stack[len(stack)-1] != "(" {
					return false
				}
				stack = stack[:len(stack)-1]
			} else if str == "]" {
				if stack[len(stack)-1] != "[" {
					return false
				}
				stack = stack[:len(stack)-1]
			} else if str == "}" {
				if stack[len(stack)-1] != "{" {
					return false
				}
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, str)
			}
		}
	}
	return len(stack) == 0
}

func longestCommonPrefix(strs []string) string {
	str := strs[0]
	for i := 1; i < len(strs); i++ {
		stringNext := strs[i]
		j := 0
		for j < len(str) && j < len(stringNext) {
			is := string(str[j])
			js := string(stringNext[j])
			if is != js {
				break
			}
			j++
		}
		if j < len(str) {
			str = str[:j]
		}
	}
	return str
}
