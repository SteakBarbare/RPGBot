package handlers

import "github.com/bwmarrin/discordgo"

// Formats the user in a readable format
func formatUser(u *discordgo.User) string {
	return u.Username + "#" + u.Discriminator
}

// Generic message format for errors
func errorMessage(title string, message string) string {
	return "❌  **" + title + "**\n" + message
}

// Generic message format for successful operations
func successMessage(title string, message string) string {
	return "✅  **" + title + "**\n" + message
}