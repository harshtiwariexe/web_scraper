package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	foundUrl := make(map[string]bool)
	seedUrl := os.Args[1:]

	channelUrl := make(chan string)
	channelFinished := make(chan bool)

	for _, url := range seedUrl {
		go crawl(url, channelUrl, channelFinished)
	}

	for c := 0; c < len(seedUrl); {
		select {
		case url := <-channelUrl:
			foundUrl[url] = true
		case <-channelFinished:
			c++
		}
	}

	fmt.Printf("\nFound %d unique URLs:\n\n", len(foundUrl))

	for url := range foundUrl {
		fmt.Println("-" + url)
	}

	close(channelUrl)

}

func crawl(url string, chUrl chan string, chFin chan bool) {
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer func() {
		chFin <- true
	}()

	b := res.Body

	defer b.Close()

	z := html.NewTokenizer(b)

	for {
		tokenT := z.Next()

		switch {

		case tokenT == html.ErrorToken:
			return

		case tokenT == html.StartTagToken:
			token := z.Token()

			isAnchorTag := token.Data == "a"

			if !isAnchorTag {
				continue
			}

			ok, url := getHref(token)

			if !ok {
				continue
			}

			hasProto := strings.Index(url, "https") == 0

			if hasProto {
				chUrl <- url
			}
		}
	}
}

func getHref(token html.Token) (ok bool, href string) {
	for _, a := range token.Attr {
		if a.Key == "href" {
			href = a.Val
			ok = true
		}
	}

	return
}
