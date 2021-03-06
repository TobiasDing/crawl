package parse

import (
	"crawl/engine"
	"regexp"
)

const regexpStr = `<a href="([^"]+)" class="tag">([^<]+)</a>`

func Content(content []byte) engine.ParseResult {
	re := regexp.MustCompile(regexpStr)
	matches := re.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:       "https://book.douban.com" + string(m[1]),
			ParseFunc: engine.NilParse,
		})
	}
	return result
	//fmt.Println(match)
}
