package handlers

import (
	"database/sql"
	"fmt"

	"github.com/SteakBarbare/RPGBot/database"
	"github.com/SteakBarbare/RPGBot/game"
	"github.com/bwmarrin/discordgo"
)

func chooseCharacterBase(s *discordgo.Session, m *discordgo.MessageCreate, selectingPlayer int, involvedPlayers []*discordgo.User) {

	var lastPlayer bool

	if selectingPlayer < len(involvedPlayers) {
		lastPlayer = false
	} else {
		lastPlayer = true
	}

	if !lastPlayer {
		s.AddHandlerOnce(chooseCharacter)
	} else {
		s.ChannelMessageSend(m.ChannelID, "All Players are ready !")
	}

}

func chooseCharacter(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "-quit" {
		s.ChannelMessageSend(m.ChannelID, "Aborting character selection")
	} else if m.Content == "-char Show" {
		//chooseCharacterBase(s, m, m.Author)
	}

	// Get the characters info from db
	charRows := database.DB.QueryRow(fmt.Sprintln("SELECT * FROM Characters WHERE player=", m.Author.ID, "AND charName=", m.Content))

	choosenChar := game.PlayerChar{}
	switch err := charRows.Scan(&choosenChar.Name, &choosenChar.Player, &choosenChar.WeaponSkill, &choosenChar.BalisticSkill, &choosenChar.Strength, &choosenChar.Endurance, &choosenChar.Agility, &choosenChar.Willpower, &choosenChar.Fellowship, &choosenChar.Hitpoints); err {
	case sql.ErrNoRows:
		s.ChannelMessageSend(m.ChannelID, "No character with this name was found, try again or check your characters with -char Show")
		//chooseCharacterBase(s, m, m.Author)
		break
	case nil:
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintln("You have selected", choosenChar.Name))
	}
}
