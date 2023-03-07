package _slice

import "fmt"

// 测试拷贝切片时遍历和内置函数性能差异
func CopyByIteration(origin []string) []string {
	ret := make([]string, len(origin))
	for i := range origin {
		ret[i] = origin[i]
	}

	return ret
}

// 测试拷贝切片时遍历和内置函数性能差异
func CopyByBuiltIn(origin []string) []string {
	ret := make([]string, len(origin))
	copy(ret, origin)
	return ret
}

func MyTest() {
	s := make([]int32, 4, 5)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	s[3] = 4
	fmt.Println(s, len(s), cap(s))
	// [:]省略默认 0:len()
	c := s[:]
	fmt.Println(c, len(c), cap(c))
	v := s[:3]
	fmt.Println(v)
	// 新切片可能不小心覆盖原切片（数组）之后的内容
	v = append(v, 9)
	fmt.Println(v)
	fmt.Println(s)
	// 扩展表达式，最后一个参数限制新切片的容量，超出时扩容，不影响原内容
	b := s[:3:3]
	b = append(b, 6)
	fmt.Println(b)
	fmt.Println(s)
}
