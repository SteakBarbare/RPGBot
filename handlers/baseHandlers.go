package handlers

import (
	"fmt"
	"strings"

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
		s.ChannelMessageSend(m.ChannelID, "Enter a name for your character, or type -quit to cancel this operation")
		s.AddHandlerOnce(NewCharacter)
		break

	// Show all the characters linked to a player
	case "-char Show":
		s.ChannelMessageSend(m.ChannelID, "Showing your characters:")
		ShowCharacters(s, m)
		break

	// DUEL BASED COMMANDS
	// Show all the characters linked to a player
	case "-duel Invite":
		s.ChannelMessageSend(m.ChannelID, "Mention the player you want to challenge or type -quit to cancel the invitation")
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
			"\n-duel Invite: Invite someone to a duel with you",
			"\nThe database will often be wiped out, so expect your characters to often disappear",
			"\n-Lambert: hahahahaha hehehehehe```"))
		break

	}
}

// Handles all checkers related reactions
func ReactionsHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	// Ignore all reactions created by the bot itself
	if r.UserID == s.State.User.ID {
		return
	}

	// Fetch some extra information about the message associated to the reaction
	m, err := s.ChannelMessage(r.ChannelID, r.MessageID)
	// Ignore reactions on messages that have an error or that have not been sent by the bot
	if err != nil || m == nil || m.Author.ID != s.State.User.ID {
		return
	}

	// Ignore messages that are not embeds with a command in the footer
	if len(m.Embeds) != 1 || m.Embeds[0].Footer == nil || m.Embeds[0].Footer.Text == "" {
		return
	}

	// Ignore reactions that haven't been set by the bot
	if !isBotReaction(s, m.Reactions, &r.Emoji) {
		return
	}

	user, err := s.User(r.UserID)
	// Ignore when sender is invalid or is a bot
	if err != nil || user == nil || user.Bot {
		return
	}

	args := strings.Split(m.Embeds[0].Footer.Text, ":")
	// Ensure valid footer command
	if len(args) != 2 {
		return
	}

	// Call the corresponding handler
	switch args[0] {
	case "duelInvite":
		duelInvitationHandler(s, r, m, user, args[1], false)
		break
	case "generalDuelInvite":
		duelInvitationHandler(s, r, m, user, args[1], true)
		break
	}
}

// Check if users reaction is one preset by the bot
func isBotReaction(s *discordgo.Session, reactions []*discordgo.MessageReactions, emoji *discordgo.Emoji) bool {
	for _, r := range reactions {
		if r.Emoji.Name == emoji.Name && r.Me {
			return true
		}
	}

	return false
}
