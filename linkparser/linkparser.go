package linkparser

import (
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href      string
	InnerText string
}

func getInnerText(node *html.Node) string {
	var sb strings.Builder

	var rec func(*html.Node)
	rec = func(n *html.Node) {
		if n.Type != html.ElementNode && n.Type != html.CommentNode {
			sb.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			rec(c)
		}
	}
	rec(node)

	return sb.String()
}

func GetLinks(document *html.Node) []Link {
	var links []Link

	var dfs func(*html.Node)
	dfs = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			// adding new link by looking at all attributes
			// attributes are listed in array, not in map
			// that's why taking certain attribute from html tag is linear time
			var l Link
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					l.Href = attr.Val
				}
			}
			l.InnerText = getInnerText(n)
			links = append(links, l)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			dfs(c)
		}
	}
	dfs(document)
	return links
}
