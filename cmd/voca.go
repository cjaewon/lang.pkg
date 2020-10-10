package cmd

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"lang.pkg/ent"
	"lang.pkg/ent/book"
	"lang.pkg/lib"
	"lang.pkg/router"
)

// Voca : Voca Application
type Voca struct {
	client *ent.Client
}

// Init : Init Voca Router
func (app Voca) Init() {
	router.Add(&router.CommandStruct{
		Match:  "add",
		Help:   "!add <단어장 코드>, <단어>, <뜻>, [예문]",
		Info:   "단어장을 생성하는 명령어 입니다.",
		PreRun: lib.Passport(app.client),
		Run:    app.addVoca,
	})
}

func (app *Voca) addVoca(s *discordgo.Session, m *discordgo.MessageCreate, cmd *router.CommandStruct) {
	args := lib.ParseContent(m.Content, cmd.Match)
	if len(args) < 3 {
		lib.CommandError(s, m, cmd)
		return
	}

	book, err := app.client.Book.
		Query().
		Where(book.BookIDEQ(args[0])).
		WithOwner().
		First(context.Background())

	if err != nil {
		log.Errorf("Failed querying book : %v", err)
		return
	}

	if book.Edges.Owner.UserID != m.Author.ID {
		s.ChannelMessageSend(m.ChannelID, "💥 자신의 단어장만 편집 가능해요. 다른 사람의 단어장을 사용하고 싶으시면 `fork` 기능을 이용해주세요")
		return
	}

	entitiy := app.client.Voca.Create().SetKey(args[1]).SetValue(args[2])
	if len(args) > 3 {
		entitiy.SetExample(args[3])
	}

	voca, err := entitiy.Save(context.Background())
	if err != nil {
		log.Errorf("Failed querying voca : %v", err)
		return
	}

	book.Update().AddVocas(voca).Save(context.Background())

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("> ✅ **%s** 를 %s에 추가하였어요", voca.Key, book.Name))
}
