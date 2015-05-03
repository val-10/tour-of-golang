package main

import "fmt"

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

var store map[string]bool

func Crawl0(url string, fetcher Fetcher, Urls chan [] string){
	body, urls, err := fetcher.Fetch(url)
	if err !=nil{
		fmt.Println(err)
	}else{
		fmt.Println("fetched:", url, body)
	}
	Urls <- urls
}

func Crawl(url string, depth int, fetcher Fetcher) {
	Urls:=make(chan[]string)
	go Crawl0(url, fetcher, Urls)
	band:=1
	store[url]=true
	for i:=0; i<depth; i++{
		for band>0{
			band--
			next:= <-Urls
			for _,url:=range next{
				if _,done:=store[url];!done{
					store [url] = true
					band++
					go Crawl0(url, fetcher, Urls)
				}}}}
	return
}

func main() {
	store=make(map[string]bool)
	Crawl("http://golang.org/", 4, fetcher)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}