package goroutine

import (
	"context"
	"net/http"
)

func server(addr string, handler http.Handler, stop <- chan struct{}) error {
	s := http.Server {
		Addr: addr,
		Handler: handler,
	}
	go func() {
		<- stop
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	go func() {
		//done <- server(stop)
	}()
	go func() {
		//done <- serveApp(stop)
	}()
	var stopped bool
	for i := 0 ; i <cap(done) ;i++ {
		if err := <- done ;err != nil {
			//..
		}
		if !stopped {
			stopped = true
			close(stop)
		}
	}
}
