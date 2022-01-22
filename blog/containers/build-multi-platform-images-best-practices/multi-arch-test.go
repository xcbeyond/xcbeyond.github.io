package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("the current platform architecture is %s.\n", runtime.GOARCH)
}
