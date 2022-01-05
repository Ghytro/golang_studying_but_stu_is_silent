package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Ghytro/golang_studying_but_stu_is_silent/linkparser"

	"golang.org/x/net/html"
)

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	htmlDocumentPath := flag.String("path", "", "A path to html document needed to parse. If the name of the document is not passed, the content will be read from stdin.")
	flag.Parse()
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
