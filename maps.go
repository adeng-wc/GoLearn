package main

import "fmt"

func main() {
	m := map[string]string{
		"name":   "name_aaaa",
		"course": "bbb",
		"city":   "beijing",
		"age":    "18",
	}

	m2 := make(map[string]int) // m2 = empty map
	var m3 map[string]int      // m3 = nil
	fmt.Println(m, m2, m3)

	fmt.Println(" Traversing map")

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println(" Getting value")
	if nameValue, ok := m["name"]; ok {
		fmt.Println(nameValue)
	} else {
		fmt.Println("key does not exist")
	}

	if nameValue, ok := m["name2"]; ok {
		fmt.Println(nameValue)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println(" Deleting  map value")

	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

	fmt.Println(lengthOfNonRepeatingSubStr("abcabc"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcfe"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcfea"))
	fmt.Println(lengthOfNonRepeatingSubStr("a"))
}
/**
	以 abcfea 为例：
	遍历的 abcfe 的时候
	a i=1 lastOccurred[a]=1
	b i=2 lastOccurred[b]=2
	c i=3 lastOccurred[c]=3
	f i=4 lastOccurred[f]=4
	e i=5 lastOccurred[e]=5

	a i=6 start=2 maxLength=6-2+1=5

------------------
	bbbb

	b	i=0	start=0	maxLength=2	lastOccurred[b]=1
	b	i=1	start=2 maxLength=2
	b	i=2

 */
// 寻找最长不含有重复字符的子串
func lengthOfNonRepeatingSubStr(s string) int {
	// map,不含重复字符的子串的起始长度,需要返回的不含重复字符的子串长度
	lastOccurred, start, maxLength := make(map[byte]int), 0, 0

	/*
		遍历每个字符:
		lastOccurred[ch] 不存在，或者 < start    ->  无需操作
		lastOccurred[ch] >= start  ->  更新 start
		更新 lastOccurred[ch] ，更新 maxLength
	  */
	for i, ch := range []byte(s) {
		// 从 map 中查字符 ch 的计数；找到，并且 >= start 下标
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		/*
			i 是遍历数组的下标
			start 是 找到元素的位置
		 */
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccurred[ch] = i
	}

	return maxLength
}
