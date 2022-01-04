package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gotutorial/linkparser"
	"gotutorial/linkparser_test"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	htmlDocumentPath := flag.String("path", "", "A path to html document needed to parse. If the name of the document is not passed, the content will be read from stdin.")
	canonizeTests := flag.Bool("canonizeTests", false, "To run this tool for test canonization, and not for usage.")
	flag.Parse()
	if *canonizeTests {
		linkparser_test.Canonize()
		return
	}
	var strDocContent string
	if *htmlDocumentPath != "" {
		docBytes, err := os.ReadFile(*htmlDocumentPath)
		strDocContent = string(docBytes)
		logError(err)
	} else {
		fmt.Fscan(os.Stdin, &strDocContent)
	}
	doc, err := html.Parse(strings.NewReader(strDocContent))
	logError(err)
	links := linkparser.GetLinks(doc)
	jsonBytes, err := json.Marshal(links)
	jsonBytes = append(append([]byte("{\"links\":"), jsonBytes...), '}')
	logError(err)
	fmt.Print(string(jsonBytes))
}
