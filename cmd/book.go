package cmd

import (
	"github.com/bwmarrin/discordgo"

	"lang.pkg/ent"
	"lang.pkg/router"
)

// Book : Book Application
type Book struct {
	client *ent.Client
}

// Init : Init Book Router
func (app Book) Init() {
	router.Add(&router.CommandStruct{
		Match: "create",
		Help:  "!create <이름>, <설명>, [공개 Y/N]",
		Info:  "단어장을 생성하는 명령어 입니다.",
		Run:   app.createBook,
	})

}

func (app *Book) createBook(s *discordgo.Session, m *discordgo.MessageCreate) {

}
