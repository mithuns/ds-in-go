package main

import (
	"fmt"

	"github.com/ds-in-go/stacks"
)

func main() {
	ms := stacks.NewMyStack()
	ms.Push("item1")
	ms.Push("item2")
	ms.Push("item3")
	fmt.Printf("\n%v", ms)
	ms.Pop()
	fmt.Printf("\n%v", ms)
	fmt.Printf("\n%v", ms.Peek())
}
