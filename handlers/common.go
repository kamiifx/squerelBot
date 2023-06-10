package handlers

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"
)

func CommonHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}
	if !strings.HasPrefix(message.Content, "sq") {
		return
	}

	args := strings.Split(message.Content, " ")[1:]
	if len(args) == 0 {
		userChannel, err := session.UserChannelCreate(message.Author.ID)
		if err != nil {
			log.Println("Error creating user channel:", err)
			return
		}

		_, err = session.ChannelMessageSend(userChannel.ID, "Listen you shit..")
		time.Sleep(2 * time.Second)
		_, err = session.ChannelMessageSend(userChannel.ID, "How about you dont use my name in vein ? Hmm ?? HOW ABOUT THAT ?")
		time.Sleep(3 * time.Second)
		_, err = session.ChannelMessageSend(userChannel.ID, "You dont know what I am capable of. YOU DONT FKN KNOW")
		time.Sleep(2 * time.Second)
		_, err = session.ChannelMessageSend(userChannel.ID, "This Squirl is built different, unless you know the command")
		time.Sleep(1 * time.Second)
		_, err = session.ChannelMessageSend(userChannel.ID, "Dont say my name.")
		if err != nil {
			log.Println("Error sending message:", err)
		}
		return
	}

	switch args[0] {
	case "pic":
		session.ChannelMessageSend(message.ChannelID, message.Author.AvatarURL("1024"))
	}
}
