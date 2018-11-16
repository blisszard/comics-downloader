package main

import (
	"fmt"
	"os"
)

func DetectComic(url string) {
	source, _ := UrlSource(url)
	splittedUrl := SplitUrl(url)

	switch source {
	case "www.comicextra.com":
		name := splittedUrl[3]
		issueNumber := splittedUrl[4]
		imageRegex := `<img[^>]+src="([^">]+)"`
		comic := NewComic(name, issueNumber, imageRegex, source)
		comic.MakeComic(url, source, splittedUrl)
	default:
		fmt.Println("This site is not supported :(")
		os.Exit(0)
	}
}