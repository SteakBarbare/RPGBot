package handlers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/SteakBarbare/RPGBot/game"

	"github.com/SteakBarbare/RPGBot/database"
	"github.com/bwmarrin/discordgo"
)

func ShowCharacters(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Get the characters info from db
	charRows, err := database.DB.Query(fmt.Sprintln("SELECT * FROM Characters WHERE player=", m.Author.ID))
	if err != nil {
		log.Fatal(err)
	}

	defer charRows.Close()

	// Show the different characters and their stats
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintln("Your characters are: "))
	for charRows.Next() {

		// Check if there is at least one character
		// if !charRows.NextResultSet() {
		// 	s.ChannelMessageSend(m.ChannelID, fmt.Sprintln("No character found"))
		// 	return
		// }

		createdCharacter := game.PlayerChar{}

		if err := charRows.Scan(&createdCharacter.Name, &createdCharacter.Player, &createdCharacter.WeaponSkill, &createdCharacter.BalisticSkill, &createdCharacter.Strength, &createdCharacter.Endurance, &createdCharacter.Agility, &createdCharacter.Willpower, &createdCharacter.Fellowship, &createdCharacter.Hitpoints); err != nil {

			log.Fatal(err)

		}

		// Send a embed message for each character showing their informations
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: fmt.Sprintln("Name: **", createdCharacter.Name, "**"),
			Description: fmt.Sprintln(
				"**WeaponSkill:** ", strconv.Itoa(createdCharacter.WeaponSkill),
				"\n**BalisticSkill:** ", strconv.Itoa(createdCharacter.BalisticSkill),
				"\n**Strength:** ", strconv.Itoa(createdCharacter.Strength),
				"\n**Endurance:** ", strconv.Itoa(createdCharacter.Endurance),
				"\n**Agility:** ", strconv.Itoa(createdCharacter.Agility),
				"\n**Willpower:** ", strconv.Itoa(createdCharacter.Willpower),
				"\n**Fellowship:** ", strconv.Itoa(createdCharacter.Fellowship),
				"\n**Hitpoints:** ", strconv.Itoa(createdCharacter.Hitpoints)),
			Color: 0x0099ff,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Player: " + m.Author.ID,
			},
		})

		if err != nil {
			s.ChannelMessageSend(m.ChannelID, errorMessage("Bot error", "Error showing characters."))
			return
		}

	}

}
