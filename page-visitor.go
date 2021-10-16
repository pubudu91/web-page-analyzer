package main

import (
	"net/http"
	"net/url"
	"regexp"

	"golang.org/x/net/html"
)

func visit(node *html.Node, info *PageInfo, host string) {
	if node.Type == html.DoctypeNode {
		info.HtmlVersion = getHTMLVersion(node)
	} else if node.Type == html.ElementNode {
		analyzeElementNode(node, info, host)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		visit(child, info, host)
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

func analyzeElementNode(node *html.Node, info *PageInfo, host string) {
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
	} else if node.Data == "a" {
		analyzeHyperlink(node, info, host)
	}
}

func analyzeHyperlink(node *html.Node, info *PageInfo, host string) {
	href := getAttribute(node.Attr, "href")

	if href == "" {
		info.Links.Inaccessible++
		info.Links.Internal++
		return
	}

	url, err := url.Parse(href)

	if err != nil {
		panic(err)
	}

	if isExternalLink(url, host) {
		info.Links.External++
	} else {
		info.Links.Internal++
	}

	if isInaccessibleLink(url) {
		info.Links.Inaccessible++
	}
}

func isExternalLink(url *url.URL, host string) bool {
	if url.Host == host {
		return false
	}

	if url.Scheme == "" {
		return false
	}

	return true
}

func isInaccessibleLink(url *url.URL) bool {
	resp, err := http.Head(url.String())
	return err != nil || resp.StatusCode < 200 || resp.StatusCode >= 300
}

func getAttribute(attr []html.Attribute, name string) string {
	href := ""

	for _, attr := range attr {
		if attr.Key == name {
			href = attr.Val
			break
		}
	}

	return href
}
