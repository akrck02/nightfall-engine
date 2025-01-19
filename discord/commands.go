package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	integerOptionMinValue          = 1.0
	dmPermission                   = false
	defaultMemberPermissions int64 = discordgo.PermissionManageServer

	commands = []*discordgo.ApplicationCommand{
		{Name: "frame", Description: "Get a frame"},
		{Name: "stop", Description: "Stop a new game"},
		{Name: "frame", Description: "Frame"},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"start": NotImplementedCommand,
		"stop":  NotImplementedCommand,
		"frame": GetFrame,
	}
)

func SyncCommands(s *discordgo.Session, guildID string) {
	existingCommands, err := s.ApplicationCommands(s.State.User.ID, guildID)
	if err != nil {
		log.Fatalf("Failed to fetch commands for guild %s: %v", guildID, err)
		return
	}

	desiredMap := make(map[string]*discordgo.ApplicationCommand)
	for _, cmd := range commands {
		desiredMap[cmd.Name] = cmd
	}

	existingMap := make(map[string]*discordgo.ApplicationCommand)
	for _, cmd := range existingCommands {
		existingMap[cmd.Name] = cmd
	}

	// Delete commands not in the desired list
	for _, cmd := range existingCommands {
		if _, found := desiredMap[cmd.Name]; !found {
			err := s.ApplicationCommandDelete(s.State.User.ID, guildID, cmd.ID)
			if err != nil {
				log.Printf(
					"Failed to delete command %s (%s) in guild %s: %v",
					cmd.Name,
					cmd.ID,
					guildID,
					err,
				)
			} else {
				log.Printf("Successfully deleted command %s (%s) in guild %s", cmd.Name, cmd.ID, guildID)
			}
		}
	}

	// Create or update existing commands
	for _, cmd := range commands {
		if existingCmd, found := existingMap[cmd.Name]; found {
			// Edit existing command
			_, err := s.ApplicationCommandEdit(s.State.User.ID, guildID, existingCmd.ID, cmd)
			if err != nil {
				log.Printf(
					"Failed to edit command %s (%s) in guild %s: %v",
					cmd.Name,
					cmd.ID,
					guildID,
					err,
				)
			} else {
				log.Printf("Successfully edited command %s (%s) in guild %s", cmd.Name, cmd.ID, guildID)
			}
		} else {
			// Create new command
			_, err := s.ApplicationCommandCreate(s.State.User.ID, guildID, cmd)
			if err != nil {
				log.Printf("Failed to create command %s in guild %s: %v", cmd.Name, guildID, err)
			} else {
				log.Printf("Successfully created command %s in guild %s", cmd.Name, guildID)
			}
		}
	}
}

func SetCommands(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func NotImplementedCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Command not implemented yet! 😘",
		},
	})
}
