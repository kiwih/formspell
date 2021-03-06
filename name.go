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
	"Ray",
	"Explosion",
	"Rain",
	"Sunbeam",
	"Blast",
	"Typhoon",
	"Void",
	"Quake",
	"Frenzy",
	"Strike",
	"Shout",
	"Song",
	"Wave",
	"Incantation",
	"Surge",
	"Curse",
	"Decimation",
	"Orb",
	"Jinx",
	"Hex",
	"Curse",
	"Slash",
	"Torch",
	"Imitation",
}

var adjectives = []string{
	"Cold",
	"Hot",
	"Blue",
	"Red",
	"Green",
	"Central",
	"Discordant",
	"Whimsical",
	"Colourful",
	"Jaunty",
	"Happy",
	"Counter",
	"Victorious",
	"Musical",
	"Flashy",
	"Holy",
	"Unholy",
	"Molten",
	"Cursed",
	"Ghostly",
	"Spectral",
	"Blazing",
	"Wavy",
	"Arcane",
	"Divine",
	"Natural",
	"Gluten-free",
	"Chaotic",
	"Aligned",
	"Moon",
	"Sun",
	"Strong",
	"Weak",
	"Honest",
	"Demonic",
	"Angelic",
	"Devilish",
	"Fiendish",
	"Gnomish",
	"Tranquil",
	"Reality",
	"False",
	"True",
	"Imitiation",
	"Serene",
	"Air",
	"Earth",
	"Fire",
	"Water",
}

var owners = []string{
	"Hammond's",
	"The Professor's",
	"The Matron's",
	"The Dean's",
}

func newRandomName() string {
	return owners[rand.Intn(len(owners))] + " " + adjectives[rand.Intn(len(adjectives))] + " " + lastwords[rand.Intn(len(lastwords))]
}
