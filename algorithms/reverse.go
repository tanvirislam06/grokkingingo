package algorithms

import (
	"fmt"
	"strings"
)

// Example function for reversing a string
func reverseString(s string) string {
	var reversed string
	length := len(s)
	for i := length - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed
}

func RunReverseString() {
	fmt.Println("Reverse String")
	fmt.Printf("%s\n", strings.Repeat("-", 100))

	listofString := []string{"123456", "hello", "AaB"}

	for _, word := range listofString {
		fmt.Printf("%s\n", reverseString(word))
	}

}
