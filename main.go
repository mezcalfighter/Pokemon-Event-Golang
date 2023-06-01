package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sess, err := discordgo.New("Bot MTExMzY1OTY3OTYxNzMzOTQ4Mg.G1oRP-.mBen2TPmZkE83V9PqMjK67FeCDM_ORCTUIBjzo")
	if err != nil {
		// do something
		log.Fatal(err)
	} else {
		sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
			if m.Author.ID == s.State.User.ID {
				return
			}
			if m.Content == "Hello" {
				s.ChannelMessageSend(m.ChannelID, "world!")
			}
		})
		_ = err
	}

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("Bot is online!")

	// blocks the channel
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
