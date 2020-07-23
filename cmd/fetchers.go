package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var (
	fetcherCount = 4
)

// FetchedHTML ...
type FetchedHTML struct {
	URL   string
	HTML  []byte
	Error error
}

func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func generator(urls []string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for _, url := range urls {
			ch <- url
		}
	}()
	return ch
}

func fetch(urlCh <-chan string) <-chan FetchedHTML {
	ch := make(chan FetchedHTML)
	go func() {
		defer close(ch)
		for url := range urlCh {
			data, err := get(url)
			ch <- FetchedHTML{URL: url, HTML: data, Error: err}
		}
	}()
	return ch
}

func join(dataChList ...<-chan FetchedHTML) <-chan FetchedHTML {
	var wg sync.WaitGroup

	ch := make(chan FetchedHTML)
	wg.Add(len(dataChList))

	for _, dataCh := range dataChList {
		go func() {
			defer wg.Done()
			for html := range dataCh {
				ch <- html
			}
		}()
	}
	go func() {
		defer close(ch)
		wg.Wait()
	}()
	return ch
}

func multiFetch(numWorkers int, inputCh <-chan string) <-chan FetchedHTML {
	var wg sync.WaitGroup

	ch := make(chan FetchedHTML, numWorkers)
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			for url := range inputCh {
				data, err := get(url)
				ch <- FetchedHTML{URL: url, HTML: data, Error: err}
			}
		}()
	}

	go func() {
		defer close(ch)
		wg.Wait()
	}()
	return ch
}

func run(urls []string) <-chan FetchedHTML {
	urlCh := generator(urls)
	htmlChs := make([]<-chan FetchedHTML, fetcherCount)
	for i := range htmlChs {
		htmlChs[i] = fetch(urlCh)
	}
	return join(htmlChs...)
}

func funRun(urls []string) <-chan FetchedHTML {
	return multiFetch(4, generator(urls))
}

func main() {
	urls := []string{
		"https://en.wikipedia.org/wiki/Ma_Chao",
		"https://en.wikipedia.org/wiki/Ma_Su",
		"https://en.wikipedia.org/wiki/Jiang_Wei",
		"https://en.wikipedia.org/wiki/Ju_Fu",
		"https://en.wikipedia.org/wiki/Liao_Hua",
		"https://en.wikipedia.org/wiki/Wei_Yan",
		"https://en.wikipedia.org/wiki/Wu_Ban",
		"https://en.wikipedia.org/wiki/Cao_Hong",
		"https://en.wikipedia.org/wiki/Cao_Ren",
		"https://en.wikipedia.org/wiki/Cao_Shuang",
		"https://en.wikipedia.org/wiki/Sima_Shi",
		"https://en.wikipedia.org/wiki/Sima_Yi",
		"https://en.wikipedia.org/wiki/Sima_Zhao",
		"https://en.wikipedia.org/wiki/Sun_Chen",
		"https://en.wikipedia.org/wiki/Zhou_Tai",
		"https://en.wikipedia.org/wiki/Zhu_Huan",
		"https://en.wikipedia.org/wiki/Zhu_Ju",
		"https://en.wikipedia.org/wiki/Zhang_Bu",
		"https://en.wikipedia.org/wiki/Zhang_Ti",
	}

	for html := range funRun(urls) {
		if html.Error != nil {
			fmt.Printf("------- Error fetching %v\n", html.Error)
		} else {
			fmt.Printf("------- Successfully fetched %s\n", html.URL)
		}
	}

	fmt.Println("------ DONE")
}
