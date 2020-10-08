package cmd

import (
	"github.com/bwmarrin/discordgo"
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

}
