package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/kiwih/formspell"
)

var (
	cr             = flag.Int("cr", 10, "The Challenge Rating of the spell")
	noDamage       = flag.Bool("nd", false, "Remove damage from spell")
	effectAddition = flag.Bool("e", false, "Add an effect to the spell")
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	spell := formspell.NewSpell(formspell.NewCR(*cr), !*noDamage, *effectAddition)
	fmt.Println(spell)
}
