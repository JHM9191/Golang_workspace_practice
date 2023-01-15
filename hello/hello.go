package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(stringutil.Reverse("Hello"))
	fmt.Println(stringutil.ToUpper("Hello"))
}
