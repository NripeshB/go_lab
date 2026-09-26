package selects

import (
	"fmt"
	"net/http"
	"time"
)

//	func Racer(a, b string) ( string) {
//		aDuration := measureResponseTime(a)
//		bDuration := measureResponseTime(b)
//
//		if aDuration < bDuration {
//			return a
//		}
//
//		return b
//	}
//
//	func measureResponseTime(url string) time.Duration {
//		start := time.Now()
//		resp, err := http.Get(url)
//		if err == nil {
//			resp.Body.Close()
//		}
//		return time.Since(start)
//	}
func Racer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil

	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(a string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		resp, err := http.Get(a)
		if err == nil {
			resp.Body.Close()
		}
		close(ch)
	}()
	return ch
}
