package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/flan0910/SafeURLBot/modules"
	"github.com/flan0910/SafeURLBot/discordmods"

	"github.com/bwmarrin/discordgo"
)

func main() {
	Token := modules.ConfigLoader().Token
	discord, err := discordgo.New(fmt.Sprintf("Bot %s", Token))
	if err != nil {
		fmt.Println(err)
	}

	discord.AddHandler(discordmods.OnMessageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer discord.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

}