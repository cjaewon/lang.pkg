package cmd

import (
	"github.com/bwmarrin/discordgo"

	"lang.pkg/router"
)

func book() {
	router.Add(&router.CommandStruct{
		Match: "create",
		Help:  "!create <이름>, <설명>, [단어장 포크 가능 Y/N]",
		Info:  "단어장을 생성하는 명령어 입니다.",
		Run:   createBook,
	})

}

func createBook(s *discordgo.Session, m *discordgo.MessageCreate) {

}
