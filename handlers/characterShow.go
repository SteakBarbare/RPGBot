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

	// Check if there is at least one character
	if !charRows.NextResultSet() {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintln("No character found"))
		return
	}

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

		if err := charRows.Scan(&charName, &player, &weaponSkill, &balisticSkill, &strength, &endurance, &agility, &willpower, &fellowship, &hitpoints); err != nil {

			log.Fatal(err)

		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintln("Name: **", charName, "**"))
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintln(
			"**WeaponSkill:** ", strconv.Itoa(weaponSkill),
			"\n**BalisticSkill:** ", strconv.Itoa(balisticSkill),
			"\n**Strength:** ", strconv.Itoa(strength),
			"\n**Endurance:** ", strconv.Itoa(endurance),
			"\n**Agility:** ", strconv.Itoa(agility),
			"\n**Willpower:** ", strconv.Itoa(willpower),
			"\n**Fellowship:** ", strconv.Itoa(fellowship),
			"\n**Hitpoints:** ", strconv.Itoa(hitpoints),
		))

	}

}
