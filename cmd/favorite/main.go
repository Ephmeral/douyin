package main

import (
	favorite "github.com/Ephmeral/douyin/kitex_gen/favorite/favoriteservice"
	"log"
)

func main() {
	svr := favorite.NewServer(new(FavoriteServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
