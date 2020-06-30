package main

import (
	"bytes"
)

func toRNA(dna string) (rna string) {

	var n = map[rune]rune{
		'C': 'G',
		'G': 'C',
		'T': 'A',
		'A': 'U',
	}
	var bb bytes.Buffer

	for _, v := range dna {
		bb.WriteRune(n[v])
	}

	return bb.String()
}
