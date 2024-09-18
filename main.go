package main

import (
	"fmt"

	"CodewarsSolutions/mumbling"
)

func main() {
	result := mumbling.Accum("abcd")
	fmt.Println(result)
}
