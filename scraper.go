package main

import (
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/mjd2021usa/tldextract"
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
	extract, _ := tldextract.New("/tmp/tld.cache", false)
	url_scheme_regex := regexp.MustCompile(`^[A-Za-z0-9+-.]+://`)
	zerowidth_regex := regexp.MustCompile("^[\u200B-\u200D\uFEFF]")
	whitespace_regex := regexp.MustCompile("\\s+")

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
			raw_url := zerowidth_regex.ReplaceAllString(maybeURL, "")
			// remove whitespaces
			raw_url = whitespace_regex.ReplaceAllString(raw_url, "")
			// remove url scheme (e.g. http:// https:// etc.)
			raw_url = url_scheme_regex.ReplaceAllString(raw_url, "")
			// remove whitespace on both ends
			raw_url = strings.TrimSpace(raw_url)
			// remove trailing slash
			raw_url = strings.TrimRight(raw_url, "/")

			extractResult := extract.Extract(raw_url)
			// check if url is valid
			if extractResult.Flag == tldextract.Domain {

				// include only subdomains, root domain and tld
				SubDomain, Domain, Tld := extractResult.SubDomain, extractResult.Domain, extractResult.Tld

				raw_url_parts := []string{SubDomain, Domain, Tld}

				n := 0
				for _, val := range raw_url_parts {
					if len(val) != 0 {
						raw_url_parts[n] = val
						n++
					}
				}
				raw_url_parts = raw_url_parts[:n]

				raw_url_ := strings.Join(raw_url_parts, ".")
				// to lower case
				strings.ToLower(raw_url_)

				// special rules for known malformed urls
				// to be deprecated when upstream fixes them
				if strings.Contains(raw_url_, "www.kenedix.com") {
					urls["www.kenedix.com"] = true
				} else if strings.Contains(raw_url_, "www.everbridgepartners.com;") {
					for _, val := range strings.Split(raw_url_, ";") {
						urls[val] = true
					}
					if strings.Contains(raw_url, ".hk") {
						urls["www.everbridgepartners.com.hk"] = true
					}
					if strings.Contains(raw_url, ".cn") {
						urls["www.everbridgepartners.com.cn"] = true
					}
				} else {
					urls[raw_url_] = true
				}
			}
		}
	}

	// get map keys as slice
	urls_ := make([]string, len(urls))
	i := 0
	for k := range urls {
		urls_[i] = k
		i++
	}

	sort.Strings(urls_) // sort alphabetically

	if len(urls_) > 0 {
		err = os.WriteFile("mas-regulated-financial-institutions.txt", []byte(strings.Join(urls_, "\n")), 0644)
	} else {
		log.Fatal("No URLs found")
	}
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}
