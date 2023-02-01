package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func goEnv(key string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	t := goEnv("DISCORD_BOT_TOKEN")
	dgo, err := discordgo.New("Bot " + t)
	if err != nil {
		fmt.Printf("error starting Discord Bot : %s", err)
	}

	// Add the message handler
	dgo.AddHandler(messageCreate)

	// Add the Guild Messages intent
	dgo.Identify.Intents = discordgo.IntentsGuildMessages

	// Connect to the gateway
	err = dgo.Open()
	if err != nil {
		fmt.Printf("error while connecting to gateway : %s", err)
		return
	}

	// Wait until CTRL + C or another signal is received
	fmt.Println("uoffin-bot is operational. Press CTRL + C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dgo.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Do not proceed if the message author is a bot
	if m.Author.Bot {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong 🏓")
		return
	}

	if m.Content == "hello" {
		s.ChannelMessageSend(m.ChannelID, "Choo choo! 🚅")
		return
	}
}
