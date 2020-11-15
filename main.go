package main

import (
	"crawl/engine"
	"crawl/parse"
)



func main() {

	url := "https://book.douban.com"

	engine.Run(engine.Request{
		Url: url,
		ParseFunc: parse.Content,

	})
}



