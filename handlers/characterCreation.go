package handlers

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/SteakBarbare/RPGBot/database"
	"github.com/SteakBarbare/RPGBot/game"
	"github.com/bwmarrin/discordgo"
)

func NewCharacter(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content != "-quit" {

		character := statsGeneration(m.Content, m.Author.ID)

		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: fmt.Sprintln("This is your character, **", m.Content, "** !\n Here are it's starting stats:"),
			Description: fmt.Sprintln(
				"**WeaponSkill:** ", strconv.Itoa(character.WeaponSkill),
				"\n**BalisticSkill:** ", strconv.Itoa(character.BalisticSkill),
				"\n**Strength:** ", strconv.Itoa(character.Strength),
				"\n**Endurance:** ", strconv.Itoa(character.Endurance),
				"\n**Agility:** ", strconv.Itoa(character.Agility),
				"\n**Willpower:** ", strconv.Itoa(character.Willpower),
				"\n**Fellowship:** ", strconv.Itoa(character.Fellowship),
				"\n**Hitpoints:** ", strconv.Itoa(character.Hitpoints)),
			Color: 0x00ff99,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Player: " + m.Author.ID,
			},
		})

		if err != nil {
			s.ChannelMessageSend(m.ChannelID, errorMessage("Bot error", "Error showing characters."))
			return
		}

		_, err = database.DB.Exec(`INSERT INTO characters(charName, player, weaponSkill, balisticSkill, strength, endurance, agility, willpower, fellowship, hitpoints) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
			m.Content, m.Author.ID, character.WeaponSkill, character.BalisticSkill, character.Strength, character.Endurance, character.Agility, character.Willpower, character.Fellowship, character.Hitpoints)

		if err != nil {
			panic(err)
		}

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
