package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebSiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds"{
		return false
	}
	return true
}


func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string,100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i<b.N;i++{
		CheckWebsites(slowStubWebsiteChecker, urls)
	}

}

func TestCheckWebsites(t *testing.T) {
	websites := []string {
		"http://google.com",
		"http://blog.ola.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":true,
		"http://blog.ola.com":true,
		"waat://furhurterwe.geds":false,
	}

	got := CheckWebsites(mockWebSiteChecker, websites)

	if !reflect.DeepEqual(want,got){
		t.Fatalf("wanted %v got %v", want,got)
	}

}