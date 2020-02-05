package main

import "fmt"

/**
反转字符串
 */
func reverse(s []int)  {
	for i, j := 0,  len(s) - 1 ;i < j; i , j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/**
判断 slice 相等
 */
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if y[i] != x[i] {
			return false
		}
	}
	return true
}

/**
slice 扩容逻辑
 */
func appendInt(x []int, y ...int) []int  {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x){
		// 仍有空间 增长
		z = x[:zlen]
	}else {
		zcap := zlen
		if zcap < 2*len(x){
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

/**
4.2.2 slice 就地修改
 */
func nonempty(strings []string) []string {
	i := 0
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			//strings[i] = s
			out = append(out,s)
			i++
		}
	}
	return out[:i]
}

func remove(slice []int, i int)  []int{
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	/**
	 4.2 slice
	 */
	// 初始化slice
	months := [...]string{1: "January", 2 : "February" , 3: "March", 4: "April", 5: "May",
		6: "June", 7 : "July", 8 : "August", 9: "September", 12: "December"}
	// 剪切
	Q2 := months[4:7]
	summer := months[6:9]

	// 打印数组
	fmt.Print(Q2)
	fmt.Print(summer)

	//数组 slice中共同元素
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("\n%s appears in both.\n", s)
			}
		}
	}

	// 超过 cap 宕机
	//fmt.Println(summer[:20])
	// 扩展
	fmt.Println(summer[:5])

	//反转数组
	a := [...]int{0,1,2,3,4,5}
	reverse(a[:])
	fmt.Println(a)

	// 向左移动两个元素
	reverse(a[:2])
	reverse(a[2:])
	reverse(a[:])
	fmt.Println(a)

	/**
	append 函数
	 */
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	// %q 单引号围绕 字符字面值
	fmt.Printf("%q\n", runes)

	var x, y []int
	for i := 0; i< 10 ;i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	/**
	4.2.2 就地修改
	 */
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)

	/**
	实现栈
	 */
	stack := []string{"a", "b", "c"}
	// push
	stack = append(stack, "d")
	// pop
	stack = stack[:len(stack) - 1]

	/**
	remove
	 */
	s := []int{5,6,7,8,9}
	remove(s, 2)
	fmt.Println(s)
}

