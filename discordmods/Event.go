package discordmods

import (
	"regexp"

	"github.com/flan0910/SafeURLBot/modules"
	"github.com/flan0910/SafeURLBot/safetest"

	"github.com/bwmarrin/discordgo"
)

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	u := m.Author
	matchText := regexp.MustCompile(`https?://([\w-]+\.)+[\w-]+(/[\w-./#?%&=]*)?`)
	if !u.Bot {
		URLs := matchText.FindAllString(m.Content, -1)
		if  len(URLs) != 0 {
			if safetest.Safebrow(URLs){
				modules.SendReply(s, m.ChannelID, "このURLは危険です！絶対にクリックしないでください！\nこのURLは危険と判断されたため削除されました。", m.Reference())
				modules.DelMsg(s, m.ChannelID, m.ID)
			}
		}
	}

}