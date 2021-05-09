package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

type SignalError struct {
	ErrorString string
}

func (e *SignalError) Error() string {
	return e.ErrorString
}

type HelloHttpHandler struct {
}

func (h *HelloHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func main() {
	s := make(chan os.Signal)
	signal.Notify(s)
	g, ctx := errgroup.WithContext(context.Background())
	stop := make(chan interface{})
	g.Go(func() error {
		select {
		case v := <-s:
			fmt.Printf("signal : %v \n\t", v)
			stop <- nil
			return &SignalError{ErrorString: v.String()}
		case <-ctx.Done():
			fmt.Println("signal done")
			return nil
		}
	})

	g.Go(func() error {
		server := &http.Server{
			Addr:    "127.0.0.1:8000",
			Handler: &HelloHttpHandler{},
		}
		go func() {
			<-stop
			fmt.Println("http server stopping")
			server.Shutdown(ctx)
			fmt.Println("http server stopped")
		}()
		fmt.Println("http server running")
		return server.ListenAndServe()

	})

	if err := g.Wait(); err != nil {
		fmt.Printf("error : %v \n\t", err)
	}

}
