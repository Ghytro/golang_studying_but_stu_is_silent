package main

import (
	"flag"
	"fmt"
	"gotutorial/website_map/graph"
	"io"
	"log"
	"net/http"
	"strings"

	"gotutorial/website_map/linkparser"

	"golang.org/x/net/html"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func debugOutput(message string) {
	if *debug {
		log.Println(message)
	}
}

type void struct{}

func dfs(g *graph.Graph, visited map[string]void, currentUrl string, maxdepth, curDepth int, sameDomain bool) {
	if curDepth > maxdepth {
		return
	}
	visited[currentUrl] = void{}
	debugOutput("Going to website: " + currentUrl)
	response, err := http.Get(currentUrl)
	checkError(err)
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Some error occurred: unable to parse response from %s\nHTTP response status: %d", currentUrl, response.StatusCode)
		return
	}
	responseBody, err := io.ReadAll(response.Body)
	checkError(err)
	doc, err := html.Parse(strings.NewReader(string(responseBody)))
	checkError(err)
	links := linkparser.GetLinks(doc)

	for _, l := range links {
		if strings.HasPrefix(l.Href, "http") && (!sameDomain || sameDomain && linkparser.GetDomain(currentUrl) == linkparser.GetDomain(l.Href)) {
			g.AddEdge(graph.Edge{currentUrl, l.Href})
			if _, ok := visited[l.Href]; !ok {
				dfs(g, visited, l.Href, maxdepth, curDepth+1, sameDomain)
			}
		}
	}
}

func GetMapGraph(rootUrl string, depth int, sameDomain bool) *graph.Graph {
	g := graph.NewGraph()
	visited := make(map[string]void)
	dfs(g, visited, rootUrl, depth, 0, sameDomain)
	return g
}

var debug *bool = new(bool)

func main() {
	websiteUrl := flag.String("url", "", "A url of a website to build a map for.")
	debug = flag.Bool("debug", false, "If you specify this flag, the program will print the debug output of the visited websites")
	depth := flag.Int("depth", 3, "Maximum depth of the search to reach different websites")
	sameDomain := flag.Bool("sameDomain", true, "Indicates if we need to build map only for the domain specified in the argument -websiteUrl")
	flag.Parse()
	if *websiteUrl == "" {
		fmt.Println("Incorrect usage of a tool. Url of a website not specified")
		return
	}
	g := GetMapGraph(*websiteUrl, *depth, *sameDomain)
	fmt.Print(g.ToGraphviz("G"))
}
