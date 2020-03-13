package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "hello ,世界"
	// 字节长度
	fmt.Println(len(s))
	// 字符长度
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s) ;  {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range s {
		fmt.Printf("%d\t%c\n", i, r)
	}
}

func basename1(s string) string  {
	for i := len(s)-1; i >= 0 ; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			s = s[:i-1]
			break
		}
	}

	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}

func comma(s string) string {
	l := len(s)
	if l <= 3 {
		return s
	}

	return comma(s[:l-3]) + "," + s[l-3:]
}

func intToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i , v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		_, _ = fmt.Fprintf(&buf, "%d", v)
	}

	buf.WriteByte(']')
	return buf.String()
}
