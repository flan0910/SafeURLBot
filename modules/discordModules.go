package modules

import (
	"fmt"
	
	"github.com/bwmarrin/discordgo"
)

func SendMessage(s *discordgo.Session, channelID string, msg string) {
	_, err := s.ChannelMessageSend(channelID, msg)
	if err != nil {
		fmt.Println("Error sending message: ", err)
	}
}

func SendReply(s *discordgo.Session, channelID string, msg string, reference *discordgo.MessageReference) {
   _, err := s.ChannelMessageSendReply(channelID, msg, reference)
   if err != nil {
   	fmt.Println("Error sending message: ", err)
   }
}

func DelMsg(s *discordgo.Session, channelID string, ID string) {
	s.ChannelMessageDelete(channelID, ID)
}