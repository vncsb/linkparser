package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/vncsb/linkparser"
	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

var filePath = flag.String("file", "examples/ex2.html", "Specifies file path")

func main() {
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Could not open the specified file!")
		return
	}

	links, err := linkparser.Parse(file)

	for i, l := range links {
		fmt.Printf("Link %v:\n", i)
		fmt.Printf("\tHref: %v\n", l.Href)
		fmt.Printf("\tText: %v\n", l.Text)
	}
}

func parseLinks(n *html.Node, links *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				*links = append(*links, Link{
					Href: a.Val,
					Text: collectText(n),
				})
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseLinks(c, links)
	}
}

func collectText(n *html.Node) string {
	var str strings.Builder
	if data := strings.TrimSpace(n.Data); n.Type == html.TextNode && data != "" {
		str.WriteString(data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		str.WriteString(collectText(c))
	}
	return str.String()
}
