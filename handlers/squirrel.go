package handlers

import (
	"github.com/bwmarrin/discordgo"
	"regexp"
	"squerelBot/api"
	"strings"
)

func containsSquirrel(message string) bool {
	lowerCaseMessage := strings.ToLower(message)
	regex := regexp.MustCompile(`\b(squirrel|squerel|squrl|squrel|squir|squel|squirtle|squirell)\b`)
	return regex.MatchString(lowerCaseMessage)
}

func containsSphyx(message string) bool {
	lowerCaseMessage := strings.ToLower(message)
	regex := regexp.MustCompile(`\b(sphonx|sphinx|spinx|sphynx|sfinx|sphynex|sphynix|spynk)\b`)
	return regex.MatchString(lowerCaseMessage)
}

func ImageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	var param string
	if containsSquirrel(message.Content) {
		param = "squirrel"
	} else if containsSphyx(message.Content) {
		param = "sphynx"
	} else {
		return
	}

	response, err := api.QueryImage(param)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, "I need to relax bit cowboy! How about you do something else chicken butt..")
		return
	}

	session.ChannelMessageSend(message.ChannelID, response)
}
