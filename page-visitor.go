package main

import (
	"net/http"
	"net/url"
	"regexp"

	"golang.org/x/net/html"
)

func visit(node *html.Node, info *PageInfo, reqURL *url.URL) {
	if node.Type == html.DoctypeNode {
		info.HtmlVersion = getHTMLVersion(node)
	} else if node.Type == html.ElementNode {
		analyzeElementNode(node, info, reqURL)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		visit(child, info, reqURL)
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

func analyzeElementNode(node *html.Node, info *PageInfo, reqURL *url.URL) {
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
		analyzeHyperlink(node, info, reqURL)
	} else if node.Data == "input" {
		analyzeInput(node, info)
	}
}

func analyzeHyperlink(node *html.Node, info *PageInfo, reqURL *url.URL) {
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

	url = reqURL.ResolveReference(url)

	if isExternalLink(url, reqURL.Host) {
		info.Links.External++
	} else {
		info.Links.Internal++
	}

	if isInaccessibleLink(url) {
		info.Links.Inaccessible++
	}
}

// Assumptions:
// - There's only one login/register forms per page
// - In a log in form, there's only one password type input field
// - In a registration form, there are two password type input fields
func analyzeInput(node *html.Node, info *PageInfo) {
	inputType := getAttribute(node.Attr, "type")
	info.HasLoginForm = info.HasLoginForm != (inputType == "password")
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
	for _, attr := range attr {
		if attr.Key == name {
			return attr.Val
		}
	}

	return ""
}
