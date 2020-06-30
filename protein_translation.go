package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	rna := "AUGUUUUCUUAAAUG"

	var codon []rune
	var protein []string
	var buffer bytes.Buffer
	var codonToAminoAcid = map[string]string{
		"AUG": "Met",
		"UUU": "Phe",
		"UUC": "Phe",
		"UUA": "Leu",
		"UUG": "Leu",
		"UCU": "Ser",
		"UCC": "Ser",
		"UCA": "Ser",
		"UCG": "Ser",
		"UGU": "Cys",
		"UGC": "Cys",
		"UAU": "Tyr",
		"UAC": "Tyr",
		"UGG": "Trp",
		"UAA": "stop",
		"UAG": "stop",
		"UGA": "stop",
	}

	codonLength := 3

	for _, v := range rna {
		codon = append(codon, v)
		if len(codon) == codonLength {
			for _, v := range codon {
				buffer.WriteRune(v)
			}

			codonString := buffer.String()
			aminoAcid := codonToAminoAcid[codonString]

			if aminoAcid == "stop" {
				fmt.Println(protein)
				os.Exit(1)
			}

			protein = append(protein, aminoAcid)

			codon = nil
			buffer.Reset()
		}
	}

	fmt.Println(rna)
	fmt.Println(protein)
}
