package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
)

const keyServerAddr = "serverAddr"

func getRoot(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	fmt.Printf("%s: got / request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "This is a website\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	ctx, cancelCtx := context.WithCancel(context.Background())

	//you define your serverOne http.Server value. This value is very similar to the HTTP server you’ve already been using, but instead of passing the address and the handler to the http.ListenAndServe function, you set them as the Addr and Handler values of the http.Server.
	serverOne := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		//BaseContext is a way to change parts of the context.Context that handler functions receive when they call the Context method of *http.Request
		BaseContext: func(listener net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, listener.Addr().String())
			return ctx
		},
	}

	serverTwo := &http.Server{
		Addr:    ":4444",
		Handler: mux,
		BaseContext: func(listener net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, listener.Addr().String())
			return ctx
		},
	}

	go func() {
		err := serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server one closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
		cancelCtx()
	}()

	go func() {
		err := serverTwo.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server two closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server two: %s\n", err)
		}
		cancelCtx()
	}()

	<-ctx.Done()

}
