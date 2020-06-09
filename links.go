package linkparser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

//Link represents a HTML <a> tag
type Link struct {
	Href string
	Text string
}

//Parse receives html content as a reader and returns the parsed links
func Parse(file io.Reader) ([]Link, error) {
	doc, err := html.Parse(file)
	if err != nil {
		return nil, err
	}

	var links []Link
	parseLinks(doc, &links)
	return links, nil
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
