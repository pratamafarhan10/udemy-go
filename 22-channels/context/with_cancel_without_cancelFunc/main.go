package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("Context:\t\t", ctx)           // context.Background
	fmt.Println("Context err:\t\t", ctx.Err()) // <nil>
	fmt.Printf("Context type:\t\t%T\n", ctx)   // *context.emptyCtx
	fmt.Println("---------------------------------------------------------------")

	ctxc, _ := context.WithCancel(ctx)

	fmt.Println("Context:\t\t", ctxc)           // context.Background
	fmt.Println("Context err:\t\t", ctxc.Err()) // <nil>
	fmt.Printf("Context type:\t\t%T\n", ctxc)   // *context.emptyCtx
	fmt.Println("---------------------------------------------------------------")

}
