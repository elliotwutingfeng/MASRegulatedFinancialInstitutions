package main

import (
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/elliotwutingfeng/go-fasttld"
)

func unique(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func main() {
	extractor, err := fasttld.New(fasttld.SuffixListParams{})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	urlSchemeRegex := regexp.MustCompile(`^[A-Za-z0-9+-.]+://`)
	zerowidthRegex := regexp.MustCompile("^[\u200B-\u200D\uFEFF]")
	whitespaceRegex := regexp.MustCompile("\\s+")

	resp, err := soup.Get("https://eservices.mas.gov.sg/fid/institution?count=0")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "class", "result-list").FindAll("a", "class", "font-resize")

	urls := make(map[string]bool)

	for _, link := range links {
		maybeURL, found := link.Attrs()["href"]

		if found && !strings.HasPrefix(maybeURL, "tel") {
			// remove zero-width spaces
			rawURL := zerowidthRegex.ReplaceAllLiteralString(maybeURL, "")
			// remove whitespaces
			rawURL = whitespaceRegex.ReplaceAllLiteralString(rawURL, "")
			// remove url scheme (e.g. http:// https:// etc.)
			rawURL = urlSchemeRegex.ReplaceAllString(rawURL, "")
			// remove whitespace on both ends
			rawURL = strings.TrimSpace(rawURL)
			// remove trailing slash
			rawURL = strings.TrimRight(rawURL, "/")
			// to lowercase
			rawURL = strings.ToLower(rawURL)
			extractResult := extractor.Extract(fasttld.URLParams{URL: rawURL})

			// check if url is valid
			if extractResult.RegisteredDomain != "" {
				// include only subdomains, root domain and tld
				SubDomain, RegisteredDomain := extractResult.SubDomain, extractResult.RegisteredDomain

				var url string
				if SubDomain != "" {
					url = SubDomain + "." + RegisteredDomain
				} else {
					url = RegisteredDomain
				}

				// special rules for known malformed urls
				// to be deprecated when upstream fixes them
				if strings.Contains(url, "www.kenedix.com") {
					urls["www.kenedix.com"] = true
				} else if strings.Contains(url, "www.everbridgepartners.com;") {
					for _, val := range strings.Split(url, ";") {
						urls[val] = true
					}
					if strings.Contains(rawURL, ".hk") {
						urls["www.everbridgepartners.com.hk"] = true
					}
					if strings.Contains(rawURL, ".cn") {
						urls["www.everbridgepartners.com.cn"] = true
					}
				} else {
					urls[url] = true
				}
			}
		}
	}

	// get map keys as slice
	sortedURLs := make([]string, len(urls))
	i := 0
	for k := range urls {
		sortedURLs[i] = k
		i++
	}

	sort.Strings(sortedURLs) // sort alphabetically

	if len(sortedURLs) > 0 {
		err = os.WriteFile("mas-regulated-financial-institutions.txt", []byte(strings.Join(sortedURLs, "\n")), 0644)
	} else {
		log.Fatal("No URLs found")
	}
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
