package main

import (
	"regexp"

	"golang.org/x/net/html"
)

func visit(node *html.Node, info *PageInfo) {
	if node.Type == html.ElementNode && node.Data == "title" {
		info.Title = node.FirstChild.Data
	} else if node.Type == html.DoctypeNode {
		info.HtmlVersion = getHTMLVersion(node)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		visit(child, info)
	}
}

func getHTMLVersion(node *html.Node) string {
	if len(node.Attr) == 0 {
		return "HTML 5"
	}

	attr := node.Attr[0]
	pattern := regexp.MustCompile("-//W3C//DTD (?P<version>[a-zA-Z0-9 .]+)//EN")
	matches := pattern.FindStringSubmatch(attr.Val)

	return matches[pattern.SubexpIndex("version")]
}
