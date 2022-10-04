package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

type contextKey string

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(context.Background(), contextKey("UserId"), 777)

	res := dbAccess(ctx)

	fmt.Fprintln(w, res)
}

func dbAccess(ctx context.Context) int {
	res := ctx.Value(contextKey("UserId")).(int)

	return res
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	fmt.Println(ctx)
	s := fmt.Sprint(ctx)

	io.WriteString(w, s)
}
