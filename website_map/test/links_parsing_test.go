package main_test

import (
	"gotutorial/website_map/linkparser"
	"testing"
)

type testSet struct {
	data   string
	answer string
}

func TestGetDomain(t *testing.T) {
	tests := []testSet{
		{"https://pkg.go.dev/strings", "pkg.go.dev"},
		{"https://www.calhoun.io/", "www.calhoun.io"},
		{"https://www.youtube.com/", "www.youtube.com"},
		{"http://vk.com", "vk.com"},
		{"https://example.com", "example.com"},
	}
	for i, test := range tests {
		a := linkparser.GetDomain(test.data)
		if a != test.answer {
			t.Fatalf("Incorrect answer for test %d: got %s but expected %s", i+1, a, test.answer)
		}
	}
}
