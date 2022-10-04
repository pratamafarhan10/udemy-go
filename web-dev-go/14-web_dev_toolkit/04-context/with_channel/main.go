package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
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

	res, err := dbAccess(ctx)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	fmt.Fprintln(w, res)
}

func dbAccess(ctx context.Context) (int, error) {
	ctx2, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ch := make(chan int)

	go func() {
		// Ridiculous long running task
		res := ctx.Value(contextKey("UserId")).(int)
		// time.Sleep(10 * time.Second)

		// Check if there is an error to stop the goroutine running
		if ctx2.Err() != nil {
			return
		}
		ch <- res
	}()

	select {
	case v := <-ch:
		return v, nil
	case <-ctx2.Done():
		return 0, ctx2.Err()
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	fmt.Println(ctx)
	s := fmt.Sprint(ctx)

	io.WriteString(w, s)
}
