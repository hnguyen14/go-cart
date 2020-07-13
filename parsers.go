package main

import (
	"bytes"
	"fmt"

	"golang.org/x/net/html"
)

// Link ...
type Link struct {
	URL string `json:"url"`
	Tag string `json:"tag"`
}

// Page ...
type Page struct {
	URL   string  `json:"url"`
	Links []*Link `json:"links"`
}

func collectText(n *html.Node, buf *bytes.Buffer) {
	if n.Type == html.TextNode {
		buf.WriteString(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectText(c, buf)
	}
}

// ParseHTML ...
func ParseHTML(content []byte) (*Page, error) {
	doc, err := html.Parse(bytes.NewReader(content))
	if err != nil {
		return nil, fmt.Errorf("HTML Parser: %w", err)
	}
	var links []*Link
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			link := Link{}
			tag := &bytes.Buffer{}
			collectText(n, tag)
			link.Tag = tag.String()
			for _, a := range n.Attr {
				if a.Key == "href" {
					link.URL = a.Val
					break
				}
			}
			links = append(links, &link)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return &Page{
		Links: links,
	}, nil
}
