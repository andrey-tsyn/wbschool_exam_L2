package main

import (
	"errors"
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

-r	—	указывает на то, что нужно рекурсивно переходить по ссылкам на сайте, чтобы скачивать страницы.
-l	—	определяет максимальную глубину вложенности страниц, которые wget должен скачать. В большинстве случаев сайты
имеют страницы с большой степенью вложенности и wget может просто «закопаться», скачивая новые страницы. Чтобы этого не
произошло можно использовать параметр -l.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// TODO Комментарии
func main() {
	recur := flag.Bool("r", false, "указывает на то, что нужно рекурсивно переходить по ссылкам на сайте, чтобы скачивать страницы.")

	depth := flag.Int("l", 5, "определяет максимальную глубину вложенности страниц, которые wget"+
		" должен скачать (по умолчанию значение равно 5, в примере мы установили 7).")

	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		println("url must be provided")
		return
	}

	if !*recur {
		*depth = 0
	}

	err := recursiveDownloadSite(args[0], 0, *depth)
	if err != nil {
		println(err.Error())
	}
}

func recursiveDownloadSite(url string, currLevel, depth int) error {
	split := strings.SplitAfter(url, "://")
	if len(split) < 2 {
		return errors.New(fmt.Sprintf("protocol must be prowided, e.x. 'https://%s' instead of '%s'", url, url))
	}
	protocol := split[0]
	urlWithoutProtocolAndQuery := split[1]
	urlWithoutProtocolAndQuery = strings.Split(urlWithoutProtocolAndQuery, "?")[0]

	err := os.MkdirAll(urlWithoutProtocolAndQuery, os.ModePerm)
	if err != nil {
		return err
	}

	fmt.Printf("Downloading '%s'...\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	file, err := os.Create(fmt.Sprintf("%s/index.html", urlWithoutProtocolAndQuery))
	if err != nil {
		return err
	}

	_, err = file.Write(content)
	if err != nil {
		return err
	}

	if depth == currLevel {
		return nil
	}

	domain := strings.Split(urlWithoutProtocolAndQuery, "/")[0]

	urls := getUrlsFromHtml(fmt.Sprintf("%s%s", protocol, domain), strings.NewReader(string(content)))

	for _, url := range urls {
		err := recursiveDownloadSite(url, currLevel+1, depth)
		if err != nil {
			return err
		}
	}

	return nil
}

func getUrlsFromHtml(mainUrl string, body io.Reader) []string {
	tokenizer := html.NewTokenizer(body)
	domain := strings.Split(mainUrl, "://")[1]
	var urls []string

TokenFor:
	for {
		tt := tokenizer.Next()

		switch tt {
		case html.ErrorToken:
			return urls
		case html.StartTagToken:
			tkn := tokenizer.Token()
			if tkn.Data == "a" {
				for _, attr := range tkn.Attr {
					if attr.Key == "href" {
						link := attr.Val
						if !strings.HasPrefix(link, "http") {
							link = mainUrl + link
						}

						if !strings.Contains(link, domain) {
							continue TokenFor
						}

						urls = append(urls, link)
					}
				}
			}
		}
	}
}
