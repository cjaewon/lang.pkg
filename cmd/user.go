package cmd

import (
	"context"
	"time"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"

	"lang.pkg/ent"
	"lang.pkg/ent/user"
	"lang.pkg/lib"
	"lang.pkg/router"
)

// User : User Application
type User struct {
	client *ent.Client
}

// Init : Init User Router
func (app User) Init() {
	router.Add(&router.CommandStruct{
		Match: "signup",
		Help:  "!signup",
		Info:  "lang.pkg 계정을 생성하는 명령어 입니다.",
		Run:   app.signUp,
	})
}

func (app *User) signUp(s *discordgo.Session, m *discordgo.MessageCreate, cmd *router.CommandStruct) {
	user, _ := app.client.User.Query().Where(user.UserIDEQ(m.Author.ID)).First(context.Background())
	if user != nil {
		s.ChannelMessageSend(m.ChannelID, "🗳️ 이미 가입하셨어요.")
		return
	}

	embedMsg, _ := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "계정 생성",
		Description: "계정을 생성하시면 lang.pkg의 다양한 기능을 이용하실 수 있어요.",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "약관동의",
				Value: "✅ 이모지를 눌러주시면 [약관](http://example.org/)에 동의하고 가입이 됩니다.",
			},
		},
	})

	s.MessageReactionAdd(embedMsg.ChannelID, embedMsg.ID, "✅")

	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancle()

	if reaction := lib.WaitForReaction(ctx, s, func(r *discordgo.MessageReactionAdd) bool {
		return r.UserID == m.Author.ID && embedMsg.ID == r.MessageID && r.Emoji.Name == "✅"
	}); reaction == nil {
		return
	}

	_, err := app.client.User.
		Create().
		SetUserID(m.Author.ID).
		SetUsername(m.Author.Username).
		SetThumbnail(m.Author.AvatarURL("256")).
		Save(context.Background())

	if err != nil {
		log.Errorf("Failed querying user : %v", err)
		return
	}

	lib.SignUpCache[m.Author.ID] = true

	s.ChannelMessageEditEmbed(m.ChannelID, embedMsg.ID, &discordgo.MessageEmbed{
		Title:       "Welcome",
		Description: "🥳 lang.pkg 계정이 생성되었습니다!",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "명령어",
				Value: "`!help` 를 입력하셔서 명령어를 확인해보세요.",
			},
		},
	})
}
