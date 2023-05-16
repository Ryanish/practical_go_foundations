package main

import "fmt"

func main() {
	banner("Go", 6)
	banner("Ryan", 10)

	s := "G*"
	fmt.Println("len", len(s))

	for i, r := range s {
		fmt.Println(i, r)
		if i == 0 {
			fmt.Printf("%c of type %T\n", r, r)
		}
	}

	b := s[0]
	fmt.Printf("%c of type %T\n", b, b)
	//byte (unint8)

	fmt.Println("g", isPalindrome("g"))
	fmt.Println("o", isPalindrome("go"))
	fmt.Println("gog", isPalindrome("gog"))

}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func banner(text string, width int) {
	padding := (width - len(text)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
