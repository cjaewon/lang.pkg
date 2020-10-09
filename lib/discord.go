package lib

import (
	"context"
	"sync"

	"github.com/bwmarrin/discordgo"
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
