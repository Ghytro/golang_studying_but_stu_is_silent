package linkparser_test

import (
	"encoding/json"
	"gotutorial/linkparser"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

const testAmount = 4

func logFatalErr(err error, t *testing.T) {
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestFirst(t *testing.T) {
	docContent, err := os.ReadFile("test1.html")
	logFatalErr(err, t)
	canonDataBytes, err := os.ReadFile("canondata.json")
	logFatalErr(err, t)
	canonData := make(map[int]string)
	err = json.Unmarshal(canonDataBytes, &canonData)
	logFatalErr(err, t)
	doc, err := html.Parse(strings.NewReader(string(docContent)))
	logFatalErr(err, t)
	links := linkparser.GetLinks(doc)
	jsonBytes, err := json.Marshal(links)
	logFatalErr(err, t)
	jsonBytes = append(append([]byte("{\"links\":"), jsonBytes...), '}')
	if string(jsonBytes) != canonData[1] {
		t.Fatalf("Wrong answer for test 1.\nExpected: %s\nBut got: %s\n", string(jsonBytes), canonData[1])
	}
}

func TestSecond(t *testing.T) {
	docContent, err := os.ReadFile("test2.html")
	logFatalErr(err, t)
	canonDataBytes, err := os.ReadFile("canondata.json")
	logFatalErr(err, t)
	var canonData map[int]string
	err = json.Unmarshal(canonDataBytes, &canonData)
	logFatalErr(err, t)
	doc, err := html.Parse(strings.NewReader(string(docContent)))
	logFatalErr(err, t)
	links := linkparser.GetLinks(doc)
	jsonBytes, err := json.Marshal(links)
	logFatalErr(err, t)
	jsonBytes = append(append([]byte("{\"links\":"), jsonBytes...), '}')
	if string(jsonBytes) != canonData[2] {
		t.Fatalf("Wrong answer for test 2.\nExpected: %s\nBut got: %s\n", string(jsonBytes), canonData[1])
	}
}

func TestThird(t *testing.T) {
	docContent, err := os.ReadFile("test3.html")
	logFatalErr(err, t)
	canonDataBytes, err := os.ReadFile("canondata.json")
	logFatalErr(err, t)
	var canonData map[int]string
	err = json.Unmarshal(canonDataBytes, &canonData)
	logFatalErr(err, t)
	doc, err := html.Parse(strings.NewReader(string(docContent)))
	logFatalErr(err, t)
	links := linkparser.GetLinks(doc)
	jsonBytes, err := json.Marshal(links)
	logFatalErr(err, t)
	jsonBytes = append(append([]byte("{\"links\":"), jsonBytes...), '}')
	if string(jsonBytes) != canonData[3] {
		t.Fatalf("Wrong answer for test 3.\nExpected: %s\nBut got: %s\n", string(jsonBytes), canonData[1])
	}
}

func TestFourth(t *testing.T) {
	docContent, err := os.ReadFile("test4.html")
	logFatalErr(err, t)
	canonDataBytes, err := os.ReadFile("canondata.json")
	logFatalErr(err, t)
	var canonData map[int]string
	err = json.Unmarshal(canonDataBytes, &canonData)
	logFatalErr(err, t)
	doc, err := html.Parse(strings.NewReader(string(docContent)))
	logFatalErr(err, t)
	links := linkparser.GetLinks(doc)
	jsonBytes, err := json.Marshal(links)
	logFatalErr(err, t)
	jsonBytes = append(append([]byte("{\"links\":"), jsonBytes...), '}')
	if string(jsonBytes) != canonData[4] {
		t.Fatalf("Wrong answer for test 4.\nExpected: %s\nBut got: %s\n", string(jsonBytes), canonData[1])
	}
}
