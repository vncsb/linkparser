package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vncsb/linkparser"
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
