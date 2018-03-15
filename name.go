package formspell

import "math/rand"

var lastwords = []string{
	"Bolt",
	"Enhancement",
	"Hailstorm",
	"Lights",
	"Humdinger",
	"Daydream",
	"Voice",
}

var adjectives = []string{
	"Cold",
	"Blue",
	"Central",
	"Discordant",
	"Whimsical",
	"Colourful",
	"Jaunty",
	"Happy",
}

func newRandomName() string {
	return "Hammond's " + adjectives[rand.Intn(len(adjectives))] + " " + lastwords[rand.Intn(len(lastwords))]
}
