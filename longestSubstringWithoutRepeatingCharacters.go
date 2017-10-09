package main

import "fmt"

func isInSubstr(data []byte, subLen, tailIndex int, value byte) int {
	newLen := 0

	for i := 0; i < subLen; i++ {
		if data[tailIndex-i] == value {
			break
		} else {
			newLen += 1
		}
	}

	return newLen
}

func lengthOfLongestSubstring(s string) int {
	byte_data := []byte(s)
	str_len := len(byte_data)
	if str_len == 0 {
		return 0
	}
	var lastRepeatSubLen = make([]int, str_len)

	lastRepeatSubLen[0] = 1
	max_value := 1
	for i := 1; i < str_len; i++ {
		t := isInSubstr(byte_data, lastRepeatSubLen[i-1], i-1, byte_data[i])
		lastRepeatSubLen[i] = t + 1
		if lastRepeatSubLen[i] > max_value {
			max_value = lastRepeatSubLen[i]
		}
	}

	return max_value
}

func main() {
	var s string = "abcabcbb"
	m := lengthOfLongestSubstring(s)
	fmt.Println(m)
}
