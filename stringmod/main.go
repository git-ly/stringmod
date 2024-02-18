package main

import (
	"fmt"

	"github.com/yangli/stringmod/quotes"
	"github.com/yangli/stringmod/strings"
)

func main() {
	o, e := strings.CountOddEven("12345")
	fmt.Println(o, e)

	fmt.Println(quotes.GetEmoji("turtle"))
}
