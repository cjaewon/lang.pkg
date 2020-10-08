package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"lang.pkg/cmd"
	"lang.pkg/router"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd.Bootstrap()
}

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("TOKEN"))

	if err != nil {
		log.Fatal("Client generating error")
	}

	discord.AddHandler(messageCreate)

	if err := discord.Open(); err != nil {
		log.Fatal("Opening connecting error")
	}

	log.Info("Discord bot is running")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sc

	discord.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !strings.HasPrefix(m.Content, "!") {
		return
	}

	router.Run(s, m)
}
