package main

import (
	"fmt"

	"github.com/akhil/discord-ping/bot"
	"github.com/akhil/discord-ping/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
