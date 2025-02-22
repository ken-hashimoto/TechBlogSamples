package main

import (
	"fmt"

	"github.com/TechBlogSamples/go-mvpkg/oldpackage"
)

func main() {
	fmt.Println(oldpackage.Ken)
	fmt.Println(oldpackage.MyFavoriteWord)

	Bob := oldpackage.Human{
		Name: "Bob",
		Age:  20,
	}
	fmt.Println(Bob)
}
