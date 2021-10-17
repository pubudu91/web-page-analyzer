package main

type PageInfo struct {
	Status       string
	HtmlVersion  string
	Title        string
	Headings     Headings
	Links        Links
	HasLoginForm bool
}

type Headings struct {
	H1 int
	H2 int
	H3 int
	H4 int
	H5 int
	H6 int
}

type Links struct {
	Internal     int
	External     int
	Inaccessible int
}
