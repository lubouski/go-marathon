package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	fmt.Println("Modules example")

	emoji.Println(":beer: Beer!!!")

	goMessage := emoji.Sprint("Internet loves :heart_eyes_cat: :joy_cat: :cat_with_wry_smile:")
	fmt.Println(goMessage)
}

