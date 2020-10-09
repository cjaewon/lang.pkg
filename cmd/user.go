package cmd

import (
	"context"
	"time"

	"github.com/bwmarrin/discordgo"
	"lang.pkg/lib"
	"lang.pkg/router"
)

func user() {
	router.Add(&router.CommandStruct{
		Match: "signup",
		Help:  "!signup",
		Info:  "lang.pkg 계정을 생성하는 명령어 입니다.",
		Run:   signUp,
	})

}

func signUp(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := discordgo.MessageEmbed{
		Title:       "계정 생성",
		Description: "계정을 생성하시면 lang.pkg의 다양한 기능을 이용하실 수 있어요.",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "약관동의",
				Value: "✅ 이모지를 눌러주시면 [약관](http://example.org/)에 동의하고 가입이 됩니다.",
			},
		},
	}

	embedMsg, _ := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
	s.MessageReactionAdd(embedMsg.ChannelID, embedMsg.ID, "✅")

	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancle()

	lib.WaitForReaction(ctx, s, func(r *discordgo.MessageReactionAdd) bool {
		if r.UserID == m.Author.ID && embedMsg.ID == r.MessageID && r.Emoji.Name == "✅" {
			return true
		}

		return false
	})
	// fmt.Println(reply.Content)
}
