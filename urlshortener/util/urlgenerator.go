package util

import (
	"fmt"
	"strings"
)

/*
	Short url generator works similarly to incrementing usual numbers.
	If you suppose, that allowedLiterals is the digits you may use in the number,
	the lastGeneratedUrlField stores the current number, which increments by calling nextShortUrl()
*/

type ShortUrlGenerator struct {
	lastGeneratedUrl      string
	allowedLiterals       string
	mappedAllowedLiterals map[byte]int // integer representation of each character, needed for counting next short url
	minShortUrlLen        int
}

func NewShortUrlGenerator(allowedLiterals string, minShortUrlLen int) *ShortUrlGenerator {
	ret := new(ShortUrlGenerator)
	ret.allowedLiterals = allowedLiterals
	ret.minShortUrlLen = minShortUrlLen
	ret.mappedAllowedLiterals = make(map[byte]int)
	for i, r := range ret.allowedLiterals {
		ret.mappedAllowedLiterals[byte(r)] = i
	}
	return ret
}

func replaceChar(s string, i int, c byte) string {
	byteArr := []byte(s)
	byteArr[i] = c
	return string(byteArr)
}

func (g ShortUrlGenerator) incChar(c byte) byte {
	if g.mappedAllowedLiterals[c] == len(g.allowedLiterals)-1 {
		return g.allowedLiterals[0]
	}
	return g.allowedLiterals[g.mappedAllowedLiterals[c]+1]
}

func (g *ShortUrlGenerator) NextShortUrl() string {
	if g.lastGeneratedUrl == "" {
		g.lastGeneratedUrl = strings.Repeat(string(g.allowedLiterals[0]), g.minShortUrlLen)
		return g.lastGeneratedUrl
	}
	g.lastGeneratedUrl = replaceChar(g.lastGeneratedUrl, len(g.lastGeneratedUrl)-1, g.incChar(g.lastGeneratedUrl[len(g.lastGeneratedUrl)-1]))
	if g.lastGeneratedUrl[len(g.lastGeneratedUrl)-1] == g.allowedLiterals[0] {
		i := len(g.lastGeneratedUrl) - 2
		for i >= 0 {
			g.lastGeneratedUrl = replaceChar(g.lastGeneratedUrl, i, g.incChar(g.lastGeneratedUrl[i]))
			if g.lastGeneratedUrl[i] != g.allowedLiterals[0] {
				break
			}
			i--
		}
		if i == -1 {
			g.lastGeneratedUrl = fmt.Sprintf("%c%s", g.allowedLiterals[0], g.lastGeneratedUrl)
		}
	}
	return g.lastGeneratedUrl
}

func (g ShortUrlGenerator) GetLastGeneratedUrl() string {
	return g.lastGeneratedUrl
}
