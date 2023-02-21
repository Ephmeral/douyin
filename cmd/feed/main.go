package main

import (
	feed "github.com/ephmeral/douyin/kitex_gen/feed/feedservice"
	"log"
)

func main() {
	svr := feed.NewServer(new(FeedServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
