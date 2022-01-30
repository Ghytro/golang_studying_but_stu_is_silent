package normalization

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type tableRow struct {
	id                               int
	phoneNumber, firstName, lastName string
}

var supportedMemoryModes []string = []string{"fullcache", "nocache"}

func isSupportedMemoryMode(memmode string) bool {
	for _, v := range supportedMemoryModes {
		if v == memmode {
			return true
		}
	}
	return false
}

func logUnsupportedMemoryMode(memmode string) {
	var errorText strings.Builder
	errorText.WriteString(fmt.Sprintf("no such memory mode: %s, supported memory modes are: ", memmode))
	for i, m := range supportedMemoryModes {
		errorText.WriteString(m)
		if i < len(supportedMemoryModes)-1 {
			errorText.WriteByte(',')
		}
	}
	LogErr(errors.New(errorText.String()))
}

func LogErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func normalizeNumber(number string) string {
	var result strings.Builder
	for _, c := range number {
		if byte(c) >= '0' && byte(c) <= '9' {
			result.WriteByte(byte(c))
		}
	}
	return result.String()
}
