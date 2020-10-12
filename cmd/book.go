package cmd

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"

	"lang.pkg/ent"
	"lang.pkg/ent/book"
	"lang.pkg/ent/user"
	"lang.pkg/lib"
	"lang.pkg/router"
)

// Book : Book Application
type Book struct {
	client *ent.Client
}

// Init : Init Book Router
func (app Book) Init() {
	router.Add(&router.CommandStruct{
		Match:  "create",
		Help:   "!create <이름>, <설명>, [공개 설정 Y/N]",
		Info:   "단어장을 생성하는 명령어 입니다.",
		PreRun: lib.Passport(app.client),
		Run:    app.createBook,
	})

	router.Add(&router.CommandStruct{
		Match:  "info",
		Help:   "!info <단어장 코드>",
		Info:   "해당 단어장의 정보를 알려주는 명령어 입니다.",
		PreRun: lib.Passport(app.client),
		Run:    app.infoBook,
	})
}

func (app *Book) createBook(s *discordgo.Session, m *discordgo.MessageCreate, cmd *router.CommandStruct) {
	args := lib.ParseContent(m.Content, cmd.Match)
	if len(args) < 2 {
		lib.CommandError(s, m, cmd)
		return
	}

	name := ""
	public := true

	if !strings.Contains(args[0], "단어장") {
		name = args[0] + " 단어장"
	} else {
		name = args[0]
	}

	if len(args) > 2 && (args[2] == "n" || args[2] == "N" || args[2] == "no" || args[2] == "NO") {
		public = false
	}

	book, err := app.client.Book.
		Create().
		SetName(name).
		SetDescription(args[1]).
		SetPublic(public).
		Save(context.Background())

	if err != nil {
		log.Errorf("Failed querying book : %v", err)
		return
	}

	book, err = app.client.Book.UpdateOneID(book.ID).SetBookID(lib.Base62Encode(book.ID)).Save(context.Background())
	if err != nil {
		log.Errorf("Failed querying book : %v", err)
		return
	}

	_, err = app.client.User.Update().Where(user.UserIDEQ(m.Author.ID)).AddBooks(book).Save(context.Background())
	if err != nil {
		log.Errorf("Failed querying user : %v", err)
		return
	}

	// TODO: 나중에 자동으로 thumnbnail, username 업데이트 기능 구현
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "📚 " + name,
		Description: "__" + args[1] + "__",
		Color:       0x70a1ff,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "공개 여부",
				Value:  map[bool]string{true: "🇴", false: "🇽"}[public],
				Inline: true,
			},
			{
				Name:   "단어 개수",
				Value:  "0개",
				Inline: true,
			},
			{
				Name:   "코드",
				Value:  "`" + *book.BookID + "`",
				Inline: true,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			IconURL: m.Author.AvatarURL("256"),
			Text:    m.Author.Username,
		},
		Timestamp: book.CreatedAt.Format(time.RFC3339),
	})
}

func (app *Book) infoBook(s *discordgo.Session, m *discordgo.MessageCreate, cmd *router.CommandStruct) {
	args := lib.ParseContent(m.Content, cmd.Match)
	if len(args) < 1 {
		lib.CommandError(s, m, cmd)
		return
	}

	book, err := app.client.Book.Query().
		Where(book.BookIDEQ((args[0]))).
		WithOwner().First(context.Background())

	if err != nil {
		log.Errorf("Failed querying book : %v", err)
		return
	}

	count, err := app.client.Voca.Query().Count(context.Background())
	if err != nil {
		log.Errorf("Failed querying book : %v", err)
		return
	}

	if book.Public == false {
		s.ChannelMessageSend(m.ChannelID, "📀 해당 단어장은 비공개로 설정되어 있어서 접근하실 수 없어요")
		return
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "📚 " + book.Name,
		Description: "__" + book.Description + "__",
		Color:       0x70a1ff,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "공개 여부",
				Value:  map[bool]string{true: "🇴", false: "🇽"}[book.Public],
				Inline: true,
			},
			{
				Name:   "단어 개수",
				Value:  strconv.Itoa(count),
				Inline: true,
			},
			{
				Name:   "코드",
				Value:  "`" + *book.BookID + "`",
				Inline: true,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			IconURL: book.Edges.Owner.Thumbnail,
			Text:    book.Edges.Owner.Username,
		},
		Timestamp: book.CreatedAt.Format(time.RFC3339),
	})
}
