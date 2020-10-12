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

	recipients := m.Mentions
	if len(recipients) == 1 {
		duelInvite(s, m, recipients[0])
	} else if len(recipients) == 0 {
		// // Ensure this is not a mistake by making sure these are the only 2 arguments
		// if len(cmd) == 1 {
		// 	sendGeneralInvite(s, m)
		// } else {
		s.ChannelMessageSend(m.ChannelID, errorMessage("Invalid Reciepient", "Ensure you are mentioning the player in the format of @<user>. Or, if you are trying to send a general invite leave the user blank."))
		return
		// }
	} else if len(recipients) > 1 {
		s.ChannelMessageSend(m.ChannelID, errorMessage("Invalid invite", "Cannot invite multiple players!"))
	} else if m.Content == "-duel Anyone" {
		s.ChannelMessageSend(m.ChannelID, "The feature isn't ready yet")
	}
}

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
			Text: "invite:" + m.Author.ID,
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
