package duels

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/SteakBarbare/RPGBot/game"
	"github.com/SteakBarbare/RPGBot/utils"
	"github.com/bwmarrin/discordgo"
)

type error interface {
	Error() string
}

func DuelController(s *discordgo.Session, channelID string, involvedPlayers []string) {

	initialSetup := duelSetup(involvedPlayers[0], involvedPlayers[1])
	fmt.Println("Challenger: ", initialSetup.Challenger)
	fmt.Println("Challenged: ", initialSetup.Challenged)

	var err error
	s.ChannelMessageSend(channelID, "Rolling Initiative...")
	initialSetup.ActiveFighter, err = rollInitiative(initialSetup, s, channelID)
	if err != nil {
		s.ChannelMessageSend(channelID, "Error when determining which character would start")
		return
	}

}

// Load Duel Infos
func duelSetup(challenger string, challenged string) *game.DuelBattle {
	initialSetup := game.DuelBattle{
		Challenger: challenger,
		Challenged: challenged,
		IsOver:     false,
		Turn:       0,
	}

	return &initialSetup
}

// Do an initiative test to determine which character will play first
func rollInitiative(duelSetup *game.DuelBattle, s *discordgo.Session, channelID string) (string, error) {
	currentDuel, err := utils.GetActiveDuel()
	if err != nil {
		return "0", errors.New("Duel not found")
	}
	currentDuelPlayers, err := utils.GetDuelPlayers(currentDuel.Id)
	if err != nil {
		return "0", errors.New("Duel data not found")
	}

	challengerChar, err := utils.GetBattleCharacterById(currentDuelPlayers.ChallengerChar)
	if err != nil {
		return "0", errors.New("Challenger character not found")
	}
	challengedChar, err := utils.GetBattleCharacterById(currentDuelPlayers.ChallengedChar)
	if err != nil {
		return "0", errors.New("Challenged character not found")
	}

	challengerInitiative := challengerChar.Agility + (rand.Intn(9) + 1)
	challengedInitiative := challengedChar.Agility + (rand.Intn(9) + 1)
	s.ChannelMessageSend(channelID, fmt.Sprintln(challengerChar.Name, " Rolled an ", challengerInitiative, " for it's initiative"))
	s.ChannelMessageSend(channelID, fmt.Sprintln(challengedChar.Name, " Rolled an ", challengedInitiative, " for it's initiative"))

	if challengerInitiative > challengedInitiative {
		s.ChannelMessageSend(channelID, fmt.Sprintln(challengerChar.Name, " will play first"))
		return duelSetup.Challenger, nil
	} else if challengedInitiative > challengerInitiative {
		s.ChannelMessageSend(channelID, fmt.Sprintln(challengedChar.Name, " will play first"))
		return duelSetup.Challenged, nil
	} else {
		s.ChannelMessageSend(channelID, "Tie ! Choosing at random who will have the initiative...")
		if rand.Intn(10) < 5 {
			s.ChannelMessageSend(channelID, fmt.Sprintln(challengerChar.Name, " will play first"))
			return duelSetup.Challenger, nil
		} else {
			s.ChannelMessageSend(channelID, fmt.Sprintln(challengedChar.Name, " will play first"))
			return duelSetup.Challenged, nil
		}
	}
}
