package main

import (
	"fmt"
	"github.com/lnzx/faker/useragent"
)

func main() {
	ua := useragent.New()
	fmt.Println(ua.Chrome())
	fmt.Println(ua.Firefox())
	fmt.Println(ua.Edge())
	fmt.Println(ua.Random())
}
