package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {

	// CHARACTER BASED COMMANDS
	// Create a new character
	case "-char New":
		s.ChannelMessageSend(m.ChannelID, "Enter a name for your character, or -quit to cancel this operation")
		s.AddHandlerOnce(NewCharacter)
		break

	// Show all the characters linked to a player
	case "-char Show":
		s.ChannelMessageSend(m.ChannelID, "Showing your characters:")
		ShowCharacters(s, m)
		break

	// DUEL BASED COMMANDS
	// Show all the characters linked to a player
	case "-duel invite":
		s.ChannelMessageSend(m.ChannelID, "Enter the name of the player you want to challenge or -quit to cancel the invitation")
		s.AddHandlerOnce(inviteCommandHandler)
		break

	// Hahahahaha hehehehehe
	case "-Lambert":
		s.ChannelMessageSend(m.ChannelID, "https://www.youtube.com/watch?v=1FswhQmILLU")
		break

	// Halp, plz, I dunno what do to with this bot :c
	case "-crrpg Help":
		s.ChannelMessageSend(m.ChannelID, "These are the different commands you can use:")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintln("```-char New: Create a new character, a name will be asked and stats are generated randomly betweend 21 & 40",
			"\n-char Show: Show all your characters and their stats",
			"\nThe database will often be wiped out, so expect your characters to often disappear```",
			"\n-Lambert: hahahahaha hehehehehe"))
		break

	}
}
