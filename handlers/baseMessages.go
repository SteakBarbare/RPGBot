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

	// Init the creation of a new character for the current user
	if m.Content == "-char New" {
		s.ChannelMessageSend(m.ChannelID, "Enter a name for your character, or -quit to cancel this operation")
		s.AddHandlerOnce(NewCharacter)
	}

	// Show the characters linked to the current user
	if m.Content == "-char Show" {
		s.ChannelMessageSend(m.ChannelID, "Showing your characters:")
		ShowCharacters(s, m)
	}

	// Show the characters linked to the current user
	if m.Content == "-crrpg Help" {
		s.ChannelMessageSend(m.ChannelID, "These are the different commands you can use:")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintln("```-char New: Create a new character, a name will be asked and stats are generated randomly betweend 21 & 40",
			"\n-char Show: Show all your characters and their stats",
			"\nThe database will often be wiped out, so expect your characters to often disappear```",
			"\n-Lambert: hahahahaha hehehehehe"))
	}

	// Hahahahaha hehehehehe
	if m.Content == "-Lambert" {
		s.ChannelMessageSend(m.ChannelID, "https://www.youtube.com/watch?v=1FswhQmILLU")
	}
}
