package game

import "github.com/bwmarrin/discordgo"

type DuelPreparation struct {
	involvedPlayers []*discordgo.User
	selectingPlayer int
	choosenChar     []string
}
