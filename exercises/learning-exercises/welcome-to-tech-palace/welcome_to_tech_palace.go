package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// createStarLine returns a string with a passed in number of asterisks
func createStarLine(numStarsPerLine int) string {
	return strings.Repeat("*", numStarsPerLine)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starLine := createStarLine(numStarsPerLine)
	return starLine + "\n" + welcomeMsg + "\n" + starLine
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}
