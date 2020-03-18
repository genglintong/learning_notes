package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

/**
1. 如果该为一个节点类型 则获取所有带有链接的节点 写入数组
2. 遍历所有子节点 递归
 */
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild ; c != nil ; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}