package router

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// CommandStruct : Command Type Struct
type CommandStruct struct {
	match string
	info  string
	help  string

	run func(s *discordgo.Session, m *discordgo.MessageCreate)
}

var commands map[string]CommandStruct

// Run : Run Command
func Run(s *discordgo.Session, m *discordgo.MessageCreate) {
	key := strings.Split(m.Content, " ")[0]

	if _, ok := commands[key]; !ok {
		return
	}

	cmd := commands[key]
	cmd.run(s, m)
}

// Bootstrap : Registe, Generate Commands
// func Bootstrap() {

// }

// Add : Add Command
func Add(key string, cmd CommandStruct) {
	if _, ok := commands[key]; ok {
		log.WithField("key", key).Error("Exits command key")
	}

	commands[key] = cmd
}
