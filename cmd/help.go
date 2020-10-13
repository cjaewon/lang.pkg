package cmd

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"lang.pkg/ent"
	"lang.pkg/lib"
	"lang.pkg/router"
)

// Help : Help Application
type Help struct {
	client *ent.Client
}

// Init : Init Help Router
func (app Help) Init() {
	router.Add(&router.CommandStruct{
		Match: "help",
		Help:  "!help [명령어]",
		Info:  "모든 명령어를 보여주거나, 해당 명령어에 대한 도움말을 제공합니다.",
		Run:   app.help,
	})
}

func (app *Help) help(s *discordgo.Session, m *discordgo.MessageCreate, cmd *router.CommandStruct) {
	args := lib.ParseContent(m.Content, cmd.Match)

	if len(args) == 0 {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       "📦 lang.pkg - 도움말",
			Description: "`!help [명령어]` 를 사용하시면 해당 명령어에 대한 정보를 보여줍니다.",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "명령어 리스트",
					Value: "`help` `signup` `create` `info` `add` `gets` `get` `remove`",
				},
			},
		})
	} else if len(args) > 0 && len(args) == 1 {
		command := router.Commands[args[0]]

		if command == nil {
			s.ChannelMessageSend(m.ChannelID, "❌ 해당 명령어를 찾을 수 없어요")
			return
		}

		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: command.Match + " 명령어 사용법",
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:  "명령어 사용법",
					Value: fmt.Sprintf("`%s`", command.Help),
				},
			},
		})
	}
}
