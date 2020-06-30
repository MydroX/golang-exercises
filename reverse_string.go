package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter text: ")
	text, _ := reader.ReadString('\n')

	textRune := make([]rune, len(text))

	index := 0
	for _, r := range text {
		textRune[index] = r
		index++
	}

	for i := 0; i < index/2; i++ {
		textRune[i], textRune[index-1-i] = textRune[index-1-i], textRune[i]
	}

	fmt.Println(string(textRune))
}
