package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("Context:\t\t", ctx)           // context.Background
	fmt.Println("Context err:\t\t", ctx.Err()) // <nil>
	fmt.Printf("Context type:\t\t%T", ctx)     // *context.emptyCtx
}
