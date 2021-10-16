package main

import (
	"regexp"

	"golang.org/x/net/html"
)

func visit(node *html.Node, info *PageInfo) {
	if node.Type == html.DoctypeNode {
		info.HtmlVersion = getHTMLVersion(node)
	} else if node.Type == html.ElementNode {
		prcoessElementNode(node, info)
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

func prcoessElementNode(node *html.Node, info *PageInfo) {
	if node.Data == "title" {
		info.Title = node.FirstChild.Data
	} else if node.Data == "h1" {
		info.Headings.H1++
	} else if node.Data == "h2" {
		info.Headings.H2++
	} else if node.Data == "h3" {
		info.Headings.H3++
	} else if node.Data == "h4" {
		info.Headings.H4++
	} else if node.Data == "h5" {
		info.Headings.H5++
	} else if node.Data == "h6" {
		info.Headings.H6++
	}
}
