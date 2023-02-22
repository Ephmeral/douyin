package main

import (
	user2 "github.com/Ephmeral/douyin/cmd/user"
	user "github.com/Ephmeral/douyin/kitex_gen/user/userservice"
	"log"
)

func main() {
	svr := user.NewServer(new(user2.UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
