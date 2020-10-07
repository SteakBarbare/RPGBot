package handlers

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/SteakBarbare/RPGBot/game"
	"github.com/bwmarrin/discordgo"
)

func NewCharacter(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content != "-quit" {

		character := statsGeneration(m.Content, m.Author.ID)

		s.ChannelMessageSend(m.ChannelID, m.Content)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintln(
			"WeaponSkill: ", strconv.Itoa(character.WeaponSkill),
			"\nBalisticSkill: ", strconv.Itoa(character.BalisticSkill),
			"\nStrength: ", strconv.Itoa(character.Strength),
			"\nEndurance: ", strconv.Itoa(character.Endurance),
			"\nAgility: ", strconv.Itoa(character.Agility),
			"\nWillpower: ", strconv.Itoa(character.Willpower),
			"\nFellowship: ", strconv.Itoa(character.Fellowship),
			"\nHitpoints: ", strconv.Itoa(character.Hitpoints),
		))

	} else {
		s.ChannelMessageSend(m.ChannelID, "Aborting character creation")
	}
}

func statsGeneration(givenName string, author string) *game.PlayerChar {
	character := game.PlayerChar{
		Name:          givenName,
		Player:        author,
		WeaponSkill:   (rand.Intn(20) + 20),
		BalisticSkill: (rand.Intn(20) + 20),
		Strength:      (rand.Intn(20) + 20),
		Endurance:     (rand.Intn(20) + 20),
		Agility:       (rand.Intn(20) + 20),
		Willpower:     (rand.Intn(20) + 20),
		Fellowship:    (rand.Intn(20) + 20),
		Hitpoints:     (rand.Intn(7) + 8)}

	return &character
}
