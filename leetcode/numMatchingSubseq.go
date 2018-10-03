package main

import "fmt"

type Seq struct {
	first int
	second int
}

func main() {
	fmt.Println("vim-go")
}

func numMatchingSubseq(S string, words []string) int {
	all := [][]Seq {}

	for i , word : range words {
		char := word[0]
		seq	:= Seq{i,1}
		append(all,strconv.Atoi(char))
	}


}
