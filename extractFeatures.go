package main

import (
	"strings"
	"unicode"
)

const (
	vovels     = "eEyYuUiIoOaAęĘóÓąĄ"
	consonants = "qwrtypsdfghjklzxcvbnmQWRTYPSDFGHJKLZXCVBNM"
	diacritics = "ęĘóÓąĄśŚłŁżŻźŹćĆńŃ"
	specials   = "`~@#$%^&*()-_=+[{]}\\|\";:,<.>/?\\'"
)

func extractFeatures(example string) extractedRow {
	var x1, x2, x3, x4, x5, x6, x7, x8 int
	for _, s := range example {
		if unicode.IsUpper(s) {
			x1 += 1
		}
		if unicode.IsLower(s) {
			x2 += 1
		}
		if unicode.IsSpace(s) {
			x3 += 1
		}
		if unicode.IsDigit(s) {
			x4 += 1
		}
		if strings.Contains(vovels, string(s)) {
			x5 += 1
		}
		if strings.Contains(consonants, string(s)) {
			x6 += 1
		}
		if strings.Contains(diacritics, string(s)) {
			x7 += 1
		}
		if strings.Contains(specials, string(s)) {
			x8 += 1
		}

	}
	return extractedRow{x1, x2, x3, x4, x5, x6, x7, x8, ""}

}

type extractedRow struct {
	UppercaseCount  int
	LowercaseCount  int
	SpaceCount      int
	DigitCount      int
	VovelsCount     int
	ConsonantsCount int
	DiacriticsCount int
	SpecialsCount   int
	Category        string
}
