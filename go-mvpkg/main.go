package main

import (
	"fmt"

	"github.com/TechBlogSamples/go-mvpkg/newpackage"
)

func main() {
	fmt.Println(oldpackage.Ken)
	fmt.Println(oldpackage.MyFavoriteWord)

	newpackage := "1" // 変更先のpackage名と同一の名前

	Bob := oldpackage.Human{
		Name: "Bob",
		Age:  20,
	}
	fmt.Println(Bob)
	fmt.Println(newpackage)
}
