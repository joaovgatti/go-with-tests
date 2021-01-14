package _select

import (
	"fmt"
	"net/http"
	"time"
)

//Non concurrent version of the application.

/*func Race(url1, url2 string) string{
	responseUrl1 := measureResponseTime(url1)
	responseUrl2 := measureResponseTime(url2)

	if responseUrl1 < responseUrl2{
		return url1
	}
	return url2
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
*/

//Concurrent version of the app

var tenSecondTimeout = 10 * time.Second

func Race(a,b string) (winner string, error error){
	return ConfigurableRacer(a,b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error){
	select {
	case <- ping(a):
		return a, nil
	case <- ping(b):
		return b, nil
	case <- time.After(timeout):
		return "",fmt.Errorf("timed out for %s and %s",a,b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func(){
		http.Get(url)
		close(ch)
	}()
	return ch
}

