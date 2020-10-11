package lib

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/bwmarrin/discordgo"
	"lang.pkg/router"
)

// MessageFilter : Filter function Type
type MessageFilter = func(m *discordgo.MessageCreate) bool

// ReactionFilter : Filter function Type
type ReactionFilter = func(r *discordgo.MessageReactionAdd) bool

// WaitForMessage : Wait for Message
func WaitForMessage(ctx context.Context, s *discordgo.Session, filter MessageFilter) *discordgo.MessageCreate {
	c := make(chan *discordgo.MessageCreate)
	once := sync.Once{}

	closer := s.AddHandler(func(_ *discordgo.Session, m *discordgo.MessageCreate) {
		if filter(m) {
			once.Do(func() {
				c <- m
				close(c)
			})
		}

	})

	defer closer()

	select {
	case <-ctx.Done():
		return nil
	case m := <-c:
		return m
	}
}

// WaitForReaction : Wait for Reaction (emoji)
func WaitForReaction(ctx context.Context, s *discordgo.Session, filter ReactionFilter) *discordgo.MessageReactionAdd {
	c := make(chan *discordgo.MessageReactionAdd)
	once := sync.Once{}

	closer := s.AddHandler(func(_ *discordgo.Session, r *discordgo.MessageReactionAdd) {
		if filter(r) {
			once.Do(func() {
				c <- r
				close(c)
			})
		}
	})

	defer closer()

	select {
	case <-ctx.Done():
		return nil
	case r := <-c:
		return r
	}
}

// MessageBotReactionRemove : Remove Bot Reaction (emoji)
func MessageBotReactionRemove(s *discordgo.Session, m *discordgo.Message, emojis ...string) {
	for _, emoji := range emojis {
		s.MessageReactionRemove(m.ChannelID, m.ID, emoji, s.State.Ready.User.ID)
	}
}

// CommandError : Send Command Error
func CommandError(s *discordgo.Session, m *discordgo.MessageCreate, cmd *router.CommandStruct) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "잘못 된 명령어",
		Description: "❌ 명령어를 잘못 사용하였어요.",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "명령어 사용법",
				Value: fmt.Sprintf("`%s`", cmd.Help),
			},
		},
	})
}

// ParseContent : Parse Message Content
func ParseContent(content string, match string) []string {
	// TODO: 따옴표로 감싸면 콤마 무시
	return MapTrim(strings.Split(strings.Replace(content, "!"+match, "", 1), ","))
}
