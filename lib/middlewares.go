package lib

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"lang.pkg/ent"
	"lang.pkg/ent/user"
)

var cache map[string]bool = map[string]bool{}

// Passport : Check User Permission
func Passport(client *ent.Client) func(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) bool {
		if val, ok := cache[m.Author.ID]; ok {
			s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+"🔒 먼저 `!signup` 명령어를 통해 계정을 만들어 주셔야 해요.")
			return val
		}

		if user, _ := client.User.Query().Where(user.IDEQ(m.Author.ID)).First(context.Background()); user == nil {
			cache[m.Author.ID] = false
			s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+"🔒 먼저 `!signup` 명령어를 통해 계정을 만들어 주셔야 해요.")

			return false
		}

		cache[m.Author.ID] = true
		return true
	}
}
