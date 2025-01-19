package discord

import (
	"bytes"
	"log"

	"github.com/bwmarrin/discordgo"

	"github.org/akrck02/nightfall/engine"
)

func JoinGame() {
}

func GetFrame(session *discordgo.Session, interactionCreate *discordgo.InteractionCreate) {
	// Create a new interaction and send images to discord chat
	currentInteraction := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Getting current game frame!",
		},
	}
	session.InteractionRespond(interactionCreate.Interaction, &currentInteraction)

	// Controls
	controls := []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					CustomID: "MoveLeft",
					Emoji:    &discordgo.ComponentEmoji{Name: "⬅️"},
					Label:    "",
					Style:    discordgo.SecondaryButton,
				},
				discordgo.Button{
					CustomID: "MoveRight",
					Emoji:    &discordgo.ComponentEmoji{Name: "➡️"},
					Label:    "",
					Style:    discordgo.SecondaryButton,
				},
				discordgo.Button{
					CustomID: "MoveUp",
					Emoji:    &discordgo.ComponentEmoji{Name: "⬆️"},
					Label:    "",
					Style:    discordgo.SecondaryButton,
				},
				discordgo.Button{
					CustomID: "MoveDown",
					Emoji:    &discordgo.ComponentEmoji{Name: "⬇️"},
					Label:    "",
					Style:    discordgo.SecondaryButton,
				},
				discordgo.Button{
					CustomID: "Interact",
					Emoji:    &discordgo.ComponentEmoji{Name: "🅰️"},
					Label:    "",
					Style:    discordgo.SecondaryButton,
				},
			},
		},
	}

	// Create edit for response
	newContent := discordgo.WebhookEdit{
		Files:      []*discordgo.File{},
		Components: &controls,
	}

	// Get current frame
	cont, err := engine.GetFrame()

	// Convert reader to discord file and add it to interaction
	discordFile := &discordgo.File{
		Name:        "Nightfall.jpg",
		Reader:      bytes.NewReader(cont),
		ContentType: "image/jpeg",
	}

	newContent.Files = append(newContent.Files, discordFile)

	// Send the images to discord chat
	_, err = session.InteractionResponseEdit(interactionCreate.Interaction, &newContent)
	if nil != err {
		log.Println(err.Error())
	}
}
