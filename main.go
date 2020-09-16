package main

import (
	"follower/config"
	"follower/util/follower"
)

func main() {
	follower.Run(follower.Config{
		Addr: config.App.AppAddr,
	})
}
