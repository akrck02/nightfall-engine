package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.org/akrck02/nightfall/configuration"
)

var session *discordgo.Session

func Start() {
	var err error
	session, err = discordgo.New("Bot " + configuration.Global.DiscordBotToken)
	if err != nil {
		panic(err)
	}

	session.AddHandler(ready)

	// We need information about guilds (which includes their channels),
	// messages and voice states.
	session.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	// Open the websocket and begin listening.
	err = session.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}
}

func Stop() {
	session.Close()
}

// Set ready function
func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateGameStatus(1, "Gaming 🎮.")
	SyncCommands(s, "")
	SetCommands(s)
}
