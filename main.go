package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"lang.pkg/cmd"
	"lang.pkg/ent"
	"lang.pkg/router"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd.Bootstrap()
}

func main() {
	client, err := ent.Open("mysql", os.Getenv("DB_URI"))
	if err != nil {
		log.Fatalf("Failed opening connection to mysql: %v", err)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Failed creating schema resources: %v", err)
	}

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
