package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// parse single file for links
func parse(dir, pathPrefix string) []Link {
	// read file
	source, err := ioutil.ReadFile(dir)
	if err != nil {
		panic(err)
	}

	// parse md
	var links []Link
	fmt.Printf("[Parsing note] %s => ", trim(dir, pathPrefix, ".md"))

	var buf bytes.Buffer
	if err := md.Convert(source, &buf); err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(&buf)
	var n int
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		target, ok := s.Attr("href")
		if !ok {
			target = "#"
		}

		target = processTarget(target)
		unesacpedTarget, _ := url.PathUnescape(target)

		source := processSource(trim(dir, pathPrefix, ".md"))
		unesacpedSource, _ := url.PathUnescape(source)

		// fmt.Printf("  '%s' => %s\n", source, target)
		links = append(links, Link{
			Source: unesacpedSource,
			Target: unesacpedTarget,
			Text:   text,
		})
		n++
	})
	fmt.Printf("found: %d links\n", n)

	return links
}
