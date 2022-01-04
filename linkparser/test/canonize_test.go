package linkparser_test

import (
	"encoding/json"
	"fmt"
	"gotutorial/linkparser"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestCanonize(t *testing.T) {
	canonData := make(map[int]string)
	for testNum := 1; testNum <= testAmount; testNum++ {
		docContent, err := os.ReadFile(fmt.Sprintf("test%d.html", testNum))
		logFatalErr(err, t)
		doc, err := html.Parse(strings.NewReader(string(docContent)))
		logFatalErr(err, t)
		links := linkparser.GetLinks(doc)
		jsonBytes, err := json.Marshal(links)
		logFatalErr(err, t)
		jsonBytes = append(append([]byte("{\"links\":"), jsonBytes...), '}')
		canonData[testNum] = string(jsonBytes)
	}
	jsonBytes, err := json.Marshal(canonData)
	logFatalErr(err, t)
	os.WriteFile("canondata.json", jsonBytes, 0644)
}
