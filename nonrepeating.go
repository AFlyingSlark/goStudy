package main

import "fmt"

func  lenthOfNonRepatingSubStr(s string) int {
	lastOccurred := make(map[rune] int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastOccurred[ch] +1
		}
		if i - start +1 > maxLength {
			maxLength = i - start +1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main()  {
	fmt.Println(
		lenthOfNonRepatingSubStr("bbbbb"))
	fmt.Println(
		lenthOfNonRepatingSubStr("assddf"))
	fmt.Println(
		lenthOfNonRepatingSubStr("qwwqrwq"))
	fmt.Println(
		lenthOfNonRepatingSubStr("aasdfasdadf"))
	fmt.Println(
		lenthOfNonRepatingSubStr("我爱大家"))
	fmt.Println(
		lenthOfNonRepatingSubStr("一一一一一"))
}