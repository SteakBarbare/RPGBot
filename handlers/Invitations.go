package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func inviteCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.Channel(m.ChannelID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, errorMessage("Bot error", "Error getting channel."))
		return
	}

	// Ensure that the command is not being sent from a dm
	if c.Type == discordgo.ChannelTypeDM {
		s.ChannelMessageSend(m.ChannelID, errorMessage("Invalid channel", "Cannot send invites from a DM"))
		return
	}

	// Check if the mention is linked to a player, or if this is a general invitation
	recipients := m.Mentions
	if len(recipients) == 1 {
		duelInvite(s, m, recipients[0])
	} else if m.Content == "-duel Anyone" {
		generalDuelInvite(s, m)
	} else if len(recipients) == 0 {
		s.ChannelMessageSend(m.ChannelID, errorMessage("Invalid Reciepient", "Ensure you are mentioning the player in the format of @<user>. Or, if you are trying to send a general invite leave the user blank."))
		return
	} else if len(recipients) > 1 {
		s.ChannelMessageSend(m.ChannelID, errorMessage("Invalid invite", "Cannot invite multiple players!"))
	}
}

// Send a DM to invite a specific player to a duel
func duelInvite(s *discordgo.Session, m *discordgo.MessageCreate, recipient *discordgo.User) {

	if m.Author.ID == recipient.ID {
		s.ChannelMessageSend(m.ChannelID, errorMessage("Invalid recipient", "Cannot play against yourself!"))
		return
	}

	if recipient.Bot {
		s.ChannelMessageSend(m.ChannelID, errorMessage("Invalid recipient", "Cannot play against bot!"))
		return
	}

	dm, err := s.UserChannelCreate(recipient.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, errorMessage("Bot error", "Error creating direct message."))
		return
	}

	invite, err := s.ChannelMessageSendEmbed(dm.ID, &discordgo.MessageEmbed{
		Title:       formatUser(m.Author) + "has challenged you to a duel",
		Description: "Click the  ✅  to accept this invitation, or the  ❌  to deny.",
		Color:       0x0099ff,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "duelInvite:" + m.Author.ID,
		},
	})

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, errorMessage("Bot error", "Error sending invite."))
		return
	}

	s.MessageReactionAdd(dm.ID, invite.ID, "✅")
	s.MessageReactionAdd(dm.ID, invite.ID, "❌")

	s.ChannelMessageSend(m.ChannelID, successMessage("Success", "Invite sent to "+formatUser(recipient)+"!"))
}

// Sends a general invite for any user in the channel to accept
func generalDuelInvite(s *discordgo.Session, m *discordgo.MessageCreate) {
	invite, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Checkers game invite from " + formatUser(m.Author),
		Description: "Click the  ✅  to accept this invitation.",
		Color:       0x0099ff,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "generalDuelInvite:" + m.Author.ID,
		},
	})

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, errorMessage("Bot error", "Error sending invite."))
		return
	}

	s.MessageReactionAdd(m.ChannelID, invite.ID, "✅")
}

// Handles all invite related reactions
func duelInvitationHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd, m *discordgo.Message, user *discordgo.User, opponentID string, general bool) {

	// If the reaction comes from the sender of the invite(This will only happen in the case of general invites)
	if r.UserID == opponentID {
		return
	}
	sender, err := s.User(opponentID)
	if err != nil || sender == nil {
		return
	}
	opponentDM, _ := s.UserChannelCreate(opponentID)

	// If accepted, send a message to confirm the invitation, then launch a duel
	if r.Emoji.Name == "✅" && (general || !hasOtherReactionsBesides("✅", m.Reactions)) {
		s.ChannelMessageEditEmbed(r.ChannelID, r.MessageID, &discordgo.MessageEmbed{
			Title:       "Invite Accepted!",
			Description: "Invite from " + formatUser(sender) + " accepted!",
			Color:       0x00ff00,
		})

		// Create a game object

		s.ChannelMessageSend(opponentDM.ID, successMessage("Game on!", formatUser(user)+" accepted your checkers invite! Wait here for them to make their move."))

		// Send a message to tell the invitation was declined otherwise
	} else if !general && r.Emoji.Name == "❌" && !hasOtherReactionsBesides("❌", m.Reactions) {
		s.ChannelMessageEditEmbed(r.ChannelID, r.MessageID, &discordgo.MessageEmbed{
			Title:       "Invite Declined",
			Description: "Invite from " + formatUser(sender) + " declined.",
			Color:       0xff0000,
		})
		s.ChannelMessageSend(opponentDM.ID, errorMessage("Invite declined", formatUser(user)+" declined your checkers game invite."))
	}
}
