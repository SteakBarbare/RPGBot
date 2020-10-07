package handlers

import "github.com/bwmarrin/discordgo"

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content == "-char New" {
		s.ChannelMessageSend(m.ChannelID, "Enter a name for your character, or -quit to cancel this operation")
		s.AddHandlerOnce(NewCharacter)
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "Lambert" {
		s.ChannelMessageSend(m.ChannelID, "https://www.youtube.com/watch?v=1FswhQmILLU")
	}
}