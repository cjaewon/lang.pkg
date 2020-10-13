package router

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// CommandStruct : Command Type Struct
type CommandStruct struct {
	Match string
	Info  string
	Help  string

	PreRun func(s *discordgo.Session, m *discordgo.MessageCreate) bool
	Run    func(s *discordgo.Session, m *discordgo.MessageCreate, cmd *CommandStruct)
}

// Commands : Command List
var Commands map[string]*CommandStruct = map[string]*CommandStruct{}

// Run : Run Command
func Run(s *discordgo.Session, m *discordgo.MessageCreate) {
	match := strings.Split(strings.TrimPrefix(m.Content, "!"), " ")[0]

	if _, ok := Commands[match]; !ok {
		return
	}

	cmd := Commands[match]

	if cmd.PreRun != nil {
		if cmd.PreRun(s, m) {
			cmd.Run(s, m, cmd)
		}

		return
	}

	cmd.Run(s, m, cmd)
}

// Add : Add Command
func Add(cmd *CommandStruct) {
	if _, ok := Commands[cmd.Match]; ok {
		log.WithField("match", cmd.Match).Error("Exits command match")
	}

	Commands[cmd.Match] = cmd
}
