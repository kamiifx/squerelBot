package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"squerelBot/handlers"
	"squerelBot/utils"
	"syscall"
)

func main() {
	token := utils.BotToken.GetValue()
	if len(token) == 0 {
		panic(utils.MissingToken.ErrorMessage())
	}

	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err.Error())
	}

	bot.AddHandler(handlers.ImageHandler)
	bot.AddHandler(handlers.CommonHandler)

	err = bot.Open()
	if err != nil {
		log.Panic(utils.DiscordConnection.ErrorMessage())
	}

	log.Print("Squerel is now up and running!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	err = bot.Close()
	if err != nil {
		return
	}
}
