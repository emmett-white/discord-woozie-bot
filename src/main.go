package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
)

var BotID string

type Config struct {
	Token  string `json:"bot_token"`
	Prefix string `json:"bot_prefix"`
}

func loadConfig(file string) (Config, error) {
	var config Config

	configFile, err := os.Open(file)
	defer configFile.Close()

	if err != nil {
		return config, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)

	return config, err
}

func main() {
	fmt.Println("Starting the application...")
	config, _ := loadConfig("./config/config.json")

	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		panic(err.Error())
	}

	u, err := dg.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID
	dg.AddHandler(handleMessages)

	err = dg.Open()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("The Woozie bot has been successfully started.")

	<-make(chan struct{})

	return
}

func handleMessages(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == s.State.User.ID {
		return
	}

	config, _ := loadConfig("./config/config.json")

	switch msg.Content {
	case (config.Prefix + "callbot"):
		s.ChannelMessageSend(msg.ChannelID, "Waddup babe?")

	case (config.Prefix + "help"):
		s.ChannelSendEmbed(msg.ChannelID, embed.NewGenericEmbed("Example", "This is an example embed!"))
	}
}
