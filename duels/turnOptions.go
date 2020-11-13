package duels

import (
	"fmt"

	"github.com/SteakBarbare/RPGBot/game"
	"github.com/SteakBarbare/RPGBot/utils"
	"github.com/bwmarrin/discordgo"
)

func FightOptionsInfo(s *discordgo.Session, m *discordgo.MessageCreate, fightSetup *game.DuelBattle) {
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: "It is your turn to play, this is what you can do",
		Description: fmt.Sprintln("```--fight Attack will do a Weaponskill test in order to hit your opponent",
			"\n-fight Dodge will do a Weaponskill test divided by 2 in order to hit, but will allow you to do a Dodge test if you get hit this turn",
			"\n-fight Flee will do an Agility test to flee this battle, if failed, you opponent get a free attack on you"),
		Color: 0x0099ff,
		// Footer: &discordgo.MessageEmbedFooter{
		// 	Text: "generalDuelInvite:" + m.Author.ID,
		// },
	})

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, utils.ErrorMessage("Bot error", "Error Showing Fight Options."))
	}

	s.AddHandlerOnce(FightTurnOptions)
}

func FightTurnOptions(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case "-fight attack":
	case "-fight dodge":
	case "-fight flee":
	}
}

func MapTurnOptions(s *discordgo.Session, m *discordgo.MessageCreate) {
}
