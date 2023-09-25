package main

import (
	"fmt"

	"github.com/admpub/sensitive"
)

func main() {
	filter := sensitive.New()
	filter.LoadWordDict("../../dict/dict.txt")
	fmt.Println(filter.Replace("xC4x", '*'))
}
