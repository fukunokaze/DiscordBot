package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"

	"github.com/auliman-a/DiscordBot/core"
	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.StringVar(&core.Token, "tw", "", "Twitter Token")
	flag.Parse()
}

func main() {
	initDiscordSession()
	// tweetText := core.GetTweet("auliman_arga")
	// fmt.Println(tweetText)
	// regInstagram := regexp.MustCompile("^https://www.instagram.com/p/+[a-zA-Z0-9]+/$")
	// if regInstagram.MatchString("https://www.instagram.com/p/CHK8tDsAqF7/") {
	// 	core.CrawlWeb("https://www.instagram.com/p/CHK8tDsAqF7/")
	// }
}

func initDiscordSession() {
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, "[IXA]") {
		command := strings.Replace(m.Content, "[IXA]", "", -1)

		switch command {
		case "gunung":
			s.ChannelMessageSend(m.ChannelID, "Kain: KUDA + panah \n Besi: INFANTRI + kuda  \n KAYU: PANAH + +infantri")
			break
		default:
			s.ChannelMessageSend(m.ChannelID, "Invalid IXA Command")

		}
	} else if strings.Contains(m.Content, "[Twitter]") {
		command := strings.Replace(m.Content, "[Twitter]", "", -1)
		tweetText := core.GetTweet(command)
		s.ChannelMessageSend(m.ChannelID, tweetText)
	}

	regInstagram := regexp.MustCompile(`^https://www\.instagram\.com/p/+.+/+.*`)
	if regInstagram.MatchString(m.Content) {
		core.CrawlWeb(m.Content)

		imgResp, _ := http.Get(core.UrlImage)

		imgFile := discordgo.File{
			Reader:      imgResp.Body,
			ContentType: "image/jpeg",
			Name:        "Test.jpeg",
		}

		msgSend := discordgo.MessageSend{
			Files:   []*discordgo.File{&imgFile},
			Content: "Instagram",
		}

		if core.UrlVideo != "" {
			vidResp, _ := http.Get(core.UrlVideo)

			vidFile := discordgo.File{
				Reader:      vidResp.Body,
				ContentType: "video/mp4",
				Name:        "Test.mp4",
			}
			msgSend.Files = append(msgSend.Files, &vidFile)
		}

		s.ChannelMessageSendComplex(m.ChannelID, &msgSend)
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// if m.Content == "test" {

	// 	resp, _ := http.Get("https://instagram.fbdo9-1.fna.fbcdn.net/v/t50.2886-16/125212583_374058524044241_8465674035460617240_n.mp4?_nc_ht=instagram.fbdo9-1.fna.fbcdn.net&_nc_cat=105&_nc_ohc=GEh1fSm6NSIAX8zjZ3F&oe=5FB47E1A&oh=f55552a4f2fef26f299b8f4624b505d7")

	// 	file := discordgo.File{
	// 		Reader:      resp.Body,
	// 		ContentType: "video/mp4",
	// 		Name:        "Test.mp4",
	// 	}

	// 	msgSend := discordgo.MessageSend{
	// 		File:    &file,
	// 		Content: "Test",
	// 	}

	// 	// video := discordgo.MessageEmbedVideo{
	// 	// 	URL: "https://instagram.fbdo9-1.fna.fbcdn.net/v/t50.2886-16/125212583_374058524044241_8465674035460617240_n.mp4?_nc_ht=instagram.fbdo9-1.fna.fbcdn.net&_nc_cat=105&_nc_ohc=GEh1fSm6NSIAX8zjZ3F&oe=5FB47E1A&oh=f55552a4f2fef26f299b8f4624b505d7",
	// 	// }

	// 	// embed := discordgo.MessageEmbed{
	// 	// 	Video: &video,
	// 	// }

	// 	s.ChannelMessageSendComplex(m.ChannelID, &msgSend)
	// }

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
