package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	return "Hello :world_map:!"
}

func main() {
	helloWorld := emoji.Sprint("Hello :world_map:!")
	fmt.Println(helloWorld)
}
