package handlers

import (
	"fmt"
	"log"
	"strconv"

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

		var (
			charName      string
			player        string
			weaponSkill   int
			balisticSkill int
			strength      int
			endurance     int
			agility       int
			willpower     int
			fellowship    int
			hitpoints     int
		)

		// Check if there is at least one character
		// if !charRows.NextResultSet() {
		// 	s.ChannelMessageSend(m.ChannelID, fmt.Sprintln("No character found"))
		// 	return
		// }

		if err := charRows.Scan(&charName, &player, &weaponSkill, &balisticSkill, &strength, &endurance, &agility, &willpower, &fellowship, &hitpoints); err != nil {

			log.Fatal(err)

		}

		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: fmt.Sprintln("Name: **", charName, "**"),
			Description: fmt.Sprintln(
				"**WeaponSkill:** ", strconv.Itoa(weaponSkill),
				"\n**BalisticSkill:** ", strconv.Itoa(balisticSkill),
				"\n**Strength:** ", strconv.Itoa(strength),
				"\n**Endurance:** ", strconv.Itoa(endurance),
				"\n**Agility:** ", strconv.Itoa(agility),
				"\n**Willpower:** ", strconv.Itoa(willpower),
				"\n**Fellowship:** ", strconv.Itoa(fellowship),
				"\n**Hitpoints:** ", strconv.Itoa(hitpoints)),
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
